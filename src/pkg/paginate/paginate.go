package paginate

import "gorm.io/gorm"

type Paginator struct {
	DB       *gorm.DB
	Page     int
	PageSize int
}

func (p *Paginator) Paginate() *gorm.DB {
	offset := (p.Page - 1) * p.PageSize
	return p.DB.Offset(offset).Limit(p.PageSize)
}
