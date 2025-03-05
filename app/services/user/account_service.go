package user

import (
	"goravel/app/models"
	"goravel/app/services/config"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

const Guard = "user"

type AccountService struct {
	configService config.BaseService
}

type LoginRequest struct {
	Account      string `form:"account"`
	Password     string `form:"password"`
	CaptchaID    string `form:"captchaID"`
	CaptchaValue string `form:"captchaValue"`
}

func (s *AccountService) Login(ctx http.Context, r LoginRequest) (user models.User, err error) {
	return
}

func (s *AccountService) User(ctx http.Context) (user models.User, err error) {
	err = facades.Auth(ctx).Guard(Guard).User(&user)
	return
}

type SaveRequest struct {
	Avatar      string `form:"avatar"`
	Name        string `form:"name"`
	Account     string `form:"account"`
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}

func (s *AccountService) Save(m models.User, r SaveRequest) error {
	return nil
}
