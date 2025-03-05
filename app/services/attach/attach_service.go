package attach

import (
	"errors"
	"fmt"

	"goravel/app/models"
	"goravel/app/services"
	"goravel/app/services/upload"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
)

type AttachService struct {
	uploadService   upload.UploadService
	categoryService CategoryService
}

type AttachSearchRequest struct {
	PageNo     int    `json:"pageNo"`
	PageSize   int    `json:"pageSize"`
	CategoryID int    `json:"categoryID"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Mime       string `json:"mime"`
}

func (s *AttachService) GetList(r AttachSearchRequest) (p services.Paginate, err error) {
	var attaches []models.Attach
	var total int64

	query := facades.Orm().Query()

	if r.CategoryID > 0 {
		query = query.Where("category_id", r.CategoryID)
	}

	if len(r.Type) > 0 {
		query = query.Where("type", r.Type)
	}

	if len(r.Name) > 0 {
		query = query.Where("name like ?", "%"+r.Name+"%")
	}

	if len(r.Mime) > 0 {
		query = query.Where("mime", r.Mime)
	}

	if err = query.Order("id desc").Paginate(r.PageNo, r.PageSize, &attaches, &total); err != nil {
		return
	}

	var resources []models.AttachResource
	for _, attache := range attaches {
		resources = append(resources, attache.Resource())
	}

	p.PageNo = r.PageNo
	p.PageSize = r.PageSize
	p.Total = total
	p.List = resources
	return
}

func (s *AttachService) GetDetail(id int) (attach models.Attach, err error) {
	err = facades.Orm().Query().FindOrFail(&attach, id)
	if err != nil {
		err = errors.New("附件不存在")
		return
	}
	return
}

func (s *AttachService) Rename(id int, name string) (err error) {
	if len(name) == 0 {
		return errors.New("文件名不能为空")
	}
	attach, err := s.GetDetail(id)
	if err != nil {
		return err
	}
	_, err = facades.Orm().Query().Model(&attach).Update("name", name)
	return
}

func (s *AttachService) Move(ids []string, categoryID int) (err error) {
	if len(ids) == 0 {
		return errors.New("附件不存在")
	}
	if categoryID > 0 {
		if _, err = s.categoryService.GetDetail(categoryID); err != nil {
			return errors.New("分组不存在")
		}
	}
	var attach models.Attach
	_, err = facades.Orm().Query().Model(&attach).Where("id", ids).Update("category_id", categoryID)
	return
}

func (s *AttachService) Delete(ids []string) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		var attaches []models.Attach
		if err := facades.Orm().Query().Where("id", ids).Get(&attaches); err != nil || len(attaches) != len(ids) {
			return errors.New("附件不存在")
		}
		for _, attach := range attaches {
			if err := s.uploadService.SetConfig(attach.Disk); err != nil {
				return err
			}
			if _, err := facades.Orm().Query().Delete(&attach); err != nil {
				return errors.New("附件记录删除失败：" + err.Error())
			}
			var driver filesystem.Driver
			if attach.Disk == "local" {
				driver = facades.Storage().Disk("public")
			} else {
				driver = facades.Storage().Disk(attach.Disk)
			}
			if !driver.Exists(attach.Path) {
				return nil
			}
			if err := driver.Delete(attach.Path); err != nil {
				message := fmt.Sprintf("附件 %s 删除失败: %v", attach.Path, err)
				facades.Log().Error(message)
				return errors.New(message)
			}
		}
		return nil
	})
}
