package commands

import (
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
)

type ResetPassword struct {
}

// Signature The name and signature of the console command.
func (receiver *ResetPassword) Signature() string {
	return "reset:password"
}

// Description The console command description.
func (receiver *ResetPassword) Description() string {
	return "Reset password by account"
}

// Extend The console command extend.
func (receiver *ResetPassword) Extend() command.Extend {
	return command.Extend{
		Category: "password",
	}
}

// Handle Execute the console command.
func (receiver *ResetPassword) Handle(ctx console.Context) (err error) {
	var admin models.Admin

	account, err := ctx.Ask("请输入要重置密码的登录账号", console.AskOption{
		Limit: 10,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("登录账号账号不能为空")
			}
			return nil
		},
	})
	if err != nil {
		return
	}
	if err = facades.Orm().Query().Where("account = ?", account).FirstOrFail(&admin); err != nil {
		ctx.Error("登录账号不存在")
		return
	}
	password, err := ctx.Secret("请输入重置密码", console.SecretOption{
		Validate: func(s string) error {
			if len(s) < 6 {
				return errors.New("密码长度至少6位")
			}
			return nil
		},
	})
	if err != nil {
		return nil
	}
	newPassword, _ := facades.Hash().Make(password)
	admin.Password = newPassword
	if err = facades.Orm().Query().Save(&admin); err != nil {
		ctx.Error("密码重置失败：" + err.Error())
		return nil
	}
	ctx.Info("密码重置成功")
	return nil
}
