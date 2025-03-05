package admin

import (
	"errors"

	"goravel/app/models"
	"goravel/app/services/captcha"
	"goravel/app/services/config"
	"goravel/app/services/permission"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
)

const Guard = "admin"

type AccountService struct {
	configService  config.BaseService
	captchaService captcha.CaptchaService
	menuService    permission.MenuService
}

type LoginRequest struct {
	Account      string `form:"account"`
	Password     string `form:"password"`
	CaptchaID    string `form:"captchaID"`
	CaptchaValue string `form:"captchaValue"`
}

func (s *AccountService) Login(ctx http.Context, r LoginRequest) (token string, err error) {
	cfg, err := s.configService.GetConfig()
	if err != nil {
		return
	}

	// 验证码校验
	if cfg.UseLoginCaptcha {
		err = s.captchaService.Verify(captcha.VerifyRequest{
			ID:    r.CaptchaID,
			Value: r.CaptchaValue,
		})
		if err != nil {
			return
		}
	}

	loginIP := ctx.Request().Ip()

	var admin models.Admin
	err = facades.Orm().Query().Where("account = ?", r.Account).FirstOrFail(&admin)
	if err != nil || !facades.Hash().Check(r.Password, admin.Password) {
		err = errors.New("登录失败，请检查账号或密码是否正确")
		return
	}

	if admin.IsEnable != models.IsEnable {
		err = errors.New("账号已被禁用，请联系系统管理员")
		return
	}

	// 设置token
	token, err = facades.Auth(ctx).Guard(Guard).LoginUsingID(&admin.ID)
	if err != nil {
		return
	}

	// 更新最近登录时间和ip
	admin.LastLoginTime = carbon.NewDateTime(carbon.Now())
	admin.LastLoginIP = loginIP
	err = facades.Orm().Query().Save(&admin)
	return
}

func (s *AccountService) GetPermissions(admin models.Admin) (treeMenus []models.TreeMenu, perms []string, err error) {
	var menus []models.Menu
	if admin.IsRoot == models.IsRoot {
		perms = []string{"*"}
		err = facades.Orm().Query().Where("is_enable", models.IsEnable).Order("sort desc,id desc").Get(&menus)
		if err != nil {
			return
		}
		treeMenus = s.menuService.ListToTree(menus, 0)
		return
	}
	if err = facades.Orm().Query().With("Roles", "is_enable", models.IsEnable).FindOrFail(&admin); err != nil {
		return
	}
	roles := admin.Roles
	if err = facades.Orm().Query().With("Menus").FindOrFail(&roles); err != nil {
		return
	}
	for _, role := range roles {
		for _, menu := range role.Menus {
			menus = append(menus, *menu)
			perms = append(perms, menu.Perm)
		}
	}
	treeMenus = s.menuService.ListToTree(menus, 0)
	return
}

func (s *AccountService) User(ctx http.Context) (admin models.Admin, err error) {
	err = facades.Auth(ctx).Guard(Guard).User(&admin)
	return
}

type SaveRequest struct {
	Avatar      string `form:"avatar"`
	Name        string `form:"name"`
	Account     string `form:"account"`
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}

func (s *AccountService) Save(m models.Admin, r SaveRequest) error {
	if r.Name != "" {
		m.Name = r.Name
	}

	if r.Account != "" {
		m.Account = r.Account
	}

	if r.OldPassword != "" && !facades.Hash().Check(r.OldPassword, m.Password) {
		return errors.New("原始密码错误")
	}

	if r.NewPassword != "" {
		password, _ := facades.Hash().Make(r.NewPassword)
		m.Password = password
	}

	m.Avatar = r.Avatar
	return facades.Orm().Query().Save(&m)
}
