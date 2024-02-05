package dao

import (
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
