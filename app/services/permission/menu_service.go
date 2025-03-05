package permission

import (
	"errors"
	"path/filepath"
	"strings"

	"goravel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type MenuService struct {
}

func (s *MenuService) GetAll() (treeMenus []models.TreeMenu, err error) {
	var menus []models.Menu
	err = facades.Orm().Query().Order("sort desc,id desc").Get(&menus)
	if err != nil || len(menus) == 0 {
		return
	}
	return s.ListToTree(menus, 0), err
}

func (s *MenuService) GetDetail(id int) (menu models.Menu, err error) {
	if err = facades.Orm().Query().FindOrFail(&menu, id); err != nil {
		err = errors.New("菜单不存在")
		return
	}
	return menu, err
}

type MenuSaveRequest struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"pid"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Sort         int    `json:"sort"`
	Url          string `json:"url"`
	Param        string `json:"param"`
	Perm         string `json:"perm"`
	Component    string `json:"component"`
	RoutePath    string `json:"routePath"`
	SelectedPath string `json:"selectedPath"`
	IsShow       int    `json:"isShow"`
	IsEnable     int    `json:"isEnable"`
}

func (s *MenuService) Save(r MenuSaveRequest) (err error) {
	var menu models.Menu

	if r.ID > 0 {
		menu, err = s.GetDetail(r.ID)
		if err != nil {
			return
		}
		menu.ID = r.ID
	}

	if len(r.Name) > 0 {
		menu.Name = r.Name
	}

	if len(r.Type) > 0 {
		menu.Type = r.Type
	}

	if len(r.Component) > 0 {
		menu.Component = r.Component
	}

	if r.Type == models.TypeButton {
		menu.IsShow = -models.IsShow
	} else {
		menu.IsShow = r.IsShow
	}

	menu.RoutePath = strings.TrimLeft(r.RoutePath, string(filepath.Separator))
	menu.SelectedPath = strings.TrimLeft(r.SelectedPath, string(filepath.Separator))
	menu.ParentID = r.ParentID
	menu.Icon = r.Icon
	menu.Param = r.Param
	menu.Url = r.Url
	menu.Perm = r.Perm
	menu.Sort = r.Sort
	menu.IsEnable = r.IsEnable

	if err = facades.Orm().Query().Save(&menu); err != nil {
		return errors.New("操作失败：" + err.Error())
	}
	return
}

func (s *MenuService) Delete(ids []string) (err error) {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		if len(ids) < 0 {
			return errors.New("请选择要删除的菜单")
		}
		var menu models.Menu
		if _, err = facades.Orm().Query().Where("id", ids).Delete(&menu); err != nil {
			return errors.New("删除失败：" + err.Error())
		}
		return nil
	})
}

func (s *MenuService) ListToTree(menus []models.Menu, parentID int) (treeMenu []models.TreeMenu) {
	if len(menus) == 0 {
		return nil
	}
	for _, menu := range menus {
		if parentID == menu.ParentID {
			treeMenu = append(treeMenu, models.TreeMenu{
				Menu:     menu,
				Children: s.ListToTree(menus, menu.ID),
			})
		}
	}
	return treeMenu
}
