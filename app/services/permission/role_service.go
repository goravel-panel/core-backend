package permission

import (
	"errors"

	"goravel/app/models"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
)

type RoleService struct {
	menuService MenuService
}

func (s *RoleService) GetAll() (roles []models.Role, err error) {
	err = facades.Orm().Query().Order("sort desc,id desc").Get(&roles)
	return
}

type RoleSearchRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	IsEnable int    `json:"isEnable"`
}

func (s *RoleService) GetList(r RoleSearchRequest) (p services.Paginate, err error) {
	var roles []models.Role
	var total int64

	query := facades.Orm().Query().With("Members", "is_enable = ?", models.IsEnable)

	if len(r.Name) > 0 {
		query = query.Where("name like ?", "%"+r.Name+"%")
	}

	if len(r.Remark) > 0 {
		query = query.Where("remark like ?", "%"+r.Remark+"%")
	}

	if r.IsEnable != 0 {
		query = query.Where("is_enable", r.IsEnable)
	}

	err = query.Order("id desc").Paginate(r.PageNo, r.PageSize, &roles, &total)

	p.PageNo = r.PageNo
	p.PageSize = r.PageSize
	p.Total = total
	p.List = roles
	return
}

func (s *RoleService) GetDetail(id int) (role models.Role, err error) {
	err = facades.Orm().Query().With("Menus", "is_enable = ?", models.IsEnable).FindOrFail(&role, id)
	if err != nil {
		err = errors.New("角色不存在")
		return
	}
	return
}

func (s *RoleService) SetAuth(role models.Role, menuIds []string) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		// 删除关联表
		if err := facades.Orm().Query().Model(&role).Association("Menus").Clear(); err != nil {
			return err
		}

		var roleMenus []models.RoleMenu
		for _, menuId := range menuIds {
			roleMenus = append(roleMenus, models.RoleMenu{
				RoleID: role.ID,
				MenuID: cast.ToInt(menuId),
			})
		}
		return facades.Orm().Query().Create(&roleMenus)
	})
}

func (s *RoleService) SetStatus(role models.Role, isEnable int) (err error) {
	_, err = facades.Orm().Query().Model(&role).Update("is_enable", isEnable)
	return
}

type RoleSaveRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	Sort     int    `json:"sort"`
	IsEnable int    `json:"isEnable"`
}

func (s *RoleService) Save(r RoleSaveRequest) (err error) {
	var role models.Role

	if r.ID > 0 {
		role, err = s.GetDetail(r.ID)
		if err != nil {
			return
		}
	}

	if len(r.Name) > 0 {
		role.Name = r.Name
	}

	if len(r.Remark) > 0 {
		role.Remark = r.Remark
	}

	role.Sort = r.Sort
	role.IsEnable = r.IsEnable

	if err = facades.Orm().Query().Save(&role); err != nil {
		return errors.New("操作失败：" + err.Error())
	}
	return nil
}

func (s *RoleService) Delete(id int) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		var role models.Role
		if err := facades.Orm().Query().With("Members").FindOrFail(&role, id); err != nil {
			return errors.New("要删除的角色不存在")
		}
		if len(role.Members) > 0 {
			return errors.New("该角色下有成员，不允许删除")
		}
		// 删除主表及关联表
		_, err := facades.Orm().Query().Select("Menus").Delete(role)
		return err
	})
}
