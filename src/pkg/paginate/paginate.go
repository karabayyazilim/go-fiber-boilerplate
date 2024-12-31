package paginate

import (
	"gorm.io/gorm"
	"sync"
)

type Paginator struct {
	DB       *gorm.DB
	Page     int
	PageSize int
	mutex    sync.RWMutex
}

func (p *Paginator) Paginate() *gorm.DB {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	dbCopy := p.DB.Session(&gorm.Session{NewDB: true})

	offset := (p.Page - 1) * p.PageSize
	return dbCopy.Offset(offset).Limit(p.PageSize)
}
