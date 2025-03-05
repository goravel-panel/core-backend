package attach

import (
	"errors"

	"goravel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type CategoryService struct {
}

func (s *CategoryService) GetAll() (categories []models.AttachCategory, err error) {
	err = facades.Orm().Query().Where("is_enable", models.IsEnable).Order("id asc").Get(&categories)
	return
}

func (s *CategoryService) GetDetail(id int) (category models.AttachCategory, err error) {
	err = facades.Orm().Query().FindOrFail(&category, id)
	if err != nil {
		err = errors.New("分组不存在")
		return
	}
	return
}

type CategorySaveRequest struct {
	ParentID int    `json:"parentID"`
	Name     string `json:"name"`
}

func (s *CategoryService) Save(r CategorySaveRequest) error {
	var category models.AttachCategory
	if len(r.Name) == 0 {
		return errors.New("分组名不能为空")
	}
	if r.ParentID > 0 {
		category.ParentID = r.ParentID
		_, err := s.GetDetail(r.ParentID)
		if err != nil {
			return errors.New("父级分组不存在")
		}
	}
	category.Name = r.Name
	category.IsEnable = models.IsEnable
	return facades.Orm().Query().Save(&category)
}

func (s *CategoryService) Rename(id int, name string) error {
	if len(name) == 0 {
		return errors.New("分组名不能为空")
	}
	category, err := s.GetDetail(id)
	if err != nil {
		return err
	}
	_, err = facades.Orm().Query().Model(&category).Update("name", name)
	return err
}

func (s *CategoryService) Delete(id int) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		var category models.AttachCategory
		if err := facades.Orm().Query().With("Attaches").FindOrFail(&category, id); err != nil {
			return errors.New("分组不存在")
		}
		if len(category.Attaches) > 0 {
			return errors.New("该分组下有文件，不允许删除")
		}
		_, err := facades.Orm().Query().Delete(&category)
		return err
	})
}
