package dao

import (
	"github.com/lgl/blog-service/internal/model"
	"github.com/lgl/blog-service/pkg/app"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(
		page,
		pageSize,
	)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: int32(id), ModifiedBy: modifiedBy},
	}
	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: int32(id)},
	}
	return tag.Delete(d.engine)
}
