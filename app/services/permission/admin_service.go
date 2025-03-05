package permission

import (
	"errors"

	"goravel/app/models"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type AdminService struct {
}

type AdminSearchRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	RoleID   int    `json:"roleID"`
	IsEnable int    `json:"isEnable"`
}

func (s *AdminService) GetList(r AdminSearchRequest) (p services.Paginate, err error) {
	var admins []models.Admin
	var total int64

	query := facades.Orm().Query().With("Roles", "is_enable = ?", models.IsEnable)

	if len(r.Account) > 0 {
		query = query.Where("account like ?", "%"+r.Account+"%")
	}

	if len(r.Name) > 0 {
		query = query.Where("name like ?", "%"+r.Name+"%")
	}

	if r.IsEnable != 0 {
		query = query.Where("is_enable", r.IsEnable)
	}

	err = query.Order("id desc").Paginate(r.PageNo, r.PageSize, &admins, &total)

	p.PageNo = r.PageNo
	p.PageSize = r.PageSize
	p.Total = total
	p.List = admins
	return
}

func (s *AdminService) GetDetail(id int) (admin models.Admin, err error) {
	err = facades.Orm().Query().With("Roles", "is_enable = ?", models.IsEnable).FindOrFail(&admin, id)
	if err != nil {
		err = errors.New("管理员不存在")
		return
	}
	return
}

func (s *AdminService) SetStatus(admin models.Admin, isEnable int) (err error) {
	if admin.IsRoot == models.IsRoot && isEnable != models.IsEnable {
		return errors.New("超级管理员不允许禁用")
	}
	_, err = facades.Orm().Query().Model(&admin).Update("is_enable", isEnable)
	return err
}

type AdminSaveRequest struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	IsEnable int    `json:"isEnable"`
	RoleIds  []int  `json:"roleIds"`
}

func (s *AdminService) Save(r AdminSaveRequest) (err error) {
	var admin models.Admin

	if r.ID > 0 {
		admin, err = s.GetDetail(r.ID)
		if err != nil {
			return
		}
	}

	if len(r.Name) > 0 {
		admin.Name = r.Name
	}

	if len(r.Account) > 0 {
		admin.Account = r.Account
	}

	if len(r.Avatar) > 0 {
		admin.Avatar = r.Avatar
	}

	if len(r.Password) > 0 {
		password, _ := facades.Hash().Make(r.Password)
		admin.Password = password
	}

	admin.IsEnable = r.IsEnable

	return facades.Orm().Transaction(func(tx orm.Query) error {
		if err = facades.Orm().Query().Save(&admin); err != nil {
			return errors.New("操作失败：" + err.Error())
		}
		if admin.IsRoot == models.IsRoot {
			return nil
		}

		// 删除关联表
		if err = facades.Orm().Query().Model(&admin).Association("Roles").Clear(); err != nil {
			return errors.New("删除失败：" + err.Error())
		}

		var adminRoles []models.AdminRole
		for _, roleId := range r.RoleIds {
			adminRoles = append(adminRoles, models.AdminRole{
				AdminID: admin.ID,
				RoleID:  roleId,
			})
		}
		return facades.Orm().Query().Create(&adminRoles)
	})
}

func (s *AdminService) Delete(id int) (err error) {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		var admin models.Admin
		if err = facades.Orm().Query().FindOrFail(&admin, id); err != nil {
			return errors.New("要删除的管理员不存在")
		}

		if admin.IsRoot == models.IsRoot {
			return errors.New("超级管理员不允许删除")
		}

		// 删除主表及关联表
		_, err = facades.Orm().Query().Select("Roles").Delete(admin)
		return err
	})
}
