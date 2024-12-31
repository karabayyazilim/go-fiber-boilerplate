package paginate

import (
	"gorm.io/gorm"
	"sync"
)

type Paginator struct {
	DB       *gorm.DB
	Page     int
	PageSize int
	mutex    sync.Mutex
}

func (p *Paginator) Paginate() *gorm.DB {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	offset := (p.Page - 1) * p.PageSize
	return p.DB.Offset(offset).Limit(p.PageSize)
}
