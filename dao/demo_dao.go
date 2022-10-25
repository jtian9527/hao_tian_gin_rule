package dao

import (
	"haotian_rule/utils"
	"gorm.io/gorm"
)

type Demo struct {
	ID       int32 `gorm:"primaryKey"`
	Name     string
	Creator  string
	Ctime    string
	Mtime    string
	IsDelete string
}

func (Demo) TableName() string {
	return "demo_demo"
}

type DemoDao struct {
	gormCluster *gorm.DB
}

var demoDao = &DemoDao{gormCluster: utils.DBCluster}

func GetDemoDao() *DemoDao {
	return demoDao
}

func (p *DemoDao) GetDemoDaoName(name string) (demos []Demo) {
	p.gormCluster.Where("name = ? AND is_delete=?", name, 0).Find(&demos)
	return
}