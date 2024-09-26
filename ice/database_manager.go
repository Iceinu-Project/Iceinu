package ice

import (
	icelogrus "github.com/Iceinu-Project/IceLogrusEnhance"
	"github.com/Iceinu-Project/Iceinu/config"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var LocalDatabase *gorm.DB  // 本地数据库连接
var RemoteDatabase *gorm.DB // 远程数据库连接

func InitLocalDatabase() {
	// 尝试和本地数据库建立连接
	gormLogger := icelogrus.NewGormLogrusLogger(log.GetLogger())
	gormLogger.LogMode(logger.Info)
	db, err := gorm.Open(sqlite.Open("iceinu.db"), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Panic("与本地SQLite数据库建立连接失败: ", err, "请检查文件系统权限！")
	}
	db.Session(&gorm.Session{Logger: gormLogger})
	// 自动迁移数据库结构
	err = db.AutoMigrate(&IceinuNodeData{}, &IceinuPluginList{})
	if err != nil {
		log.Panic("自动迁移数据库结构失败: ", err)
	}
	// 查询是否为第一次进行初始化节点数据
	var count int64
	db.Model(&IceinuNodeData{}).Count(&count)
	if count == 0 {
		log.Infof("检测到数据库中没有节点数据，正在初始化节点数据...")
		// 如果是第一次初始化，插入一个节点数据
		SelfNodeId = GenerateNodeId()
		db.Create(&IceinuNodeData{
			NodeId:       SelfNodeId,
			AdapterModel: "satori",
		})
	} else {
		// 否则读取节点数据
		var nodeData IceinuNodeData
		db.First(&nodeData)
		// 更新全局NodeId
		SetSelfNodeId(nodeData.NodeId)
	}
	// 设置数据库连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("获取数据库连接池失败: ", err)
	}
	sqlDB.SetMaxIdleConns(config.IceConf.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.IceConf.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.IceConf.Database.ConnMaxLifetime) * time.Minute)

	// 设置本地数据库连接
	LocalDatabase = db
}
