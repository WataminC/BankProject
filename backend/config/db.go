package config

import (
	"bank/global"
	"bank/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDB() error {
	err := global.DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{}, &models.LoanRequest{}, &models.Loan{})
	if err != nil {
		return err
	}

	return nil
}

func InitDB() error {
	// 获取数据库连接字符串
	dsn := AppConfig.DataBase.Dsn

	// 使用 GORM v2 初始化数据库连接
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// 获取底层数据库连接对象
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get the database object: %v", err)
	}

	// 配置数据库连接池
	sqlDB.SetMaxIdleConns(AppConfig.DataBase.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.DataBase.MaxOpennConns)

	var dbName string
	err = DB.Raw("SELECT current_database()").Scan(&dbName).Error
	if err != nil {
		log.Fatalf("Failed to fetch current database name: %v", err)
	}
	fmt.Printf("Connected to database: %s\n", dbName)

	// 打印成功连接日志
	fmt.Println("Successfully connected to the database")

	global.DB = DB

	if err := MigrateDB(); err != nil {
		log.Fatalf("Error whenn migrating database: %v", err)
	}

	return nil
}
