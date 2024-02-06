package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hao6v/internal/model"
	"hao6v/pkg/global"
)

type PageDao struct {
}

func NewPageDao() *PageDao {
	return &PageDao{}
}

func (p *PageDao) FirstOrCreate(page *model.Page) (*model.Page, int64, error) {
	ret := global.DB.Where(model.Page{AbsoluteUrl: page.AbsoluteUrl}).FirstOrCreate(page)
	return page, ret.RowsAffected, ret.Error
}

func (p *PageDao) GetRawUrl() (page *model.Page, err error) {
	// 拿到一个等待处理的链接
	r := global.DB.Where(&model.Page{Status: 0}).Last(&page)

	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			fmt.Println("记录未找到")
		}
		return nil, r.Error
	}

	return page, nil
}
