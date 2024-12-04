package postgres

import (
	"fmt"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/global"
	"project/models"
)

const port = 5432

// const api = ""
// const usr = ""
// const pwd = ""

var DB *gorm.DB

// todo
func init() {
	var err error

	// 链接 postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		global.PostgresAPI, global.PostgresUsr, global.PostgresPwd,
		global.PostgresDB, port)
	DB, err = gorm.Open(driver.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}, //自动迁移时，禁用外键约束,不禁用
	)

	if err != nil {
		// 出错
		fmt.Println(err)
		return
	}

	// 迁移模型
	err = DB.AutoMigrate(
		&models.Account{}, // 账户
		//&models.Operator{}, // 管理员
		&models.Order{}, // 订单
	)

	// 创建管理员账号
	for i := 0; i < len(global.OperatorList); i++ {
		operator := global.OperatorList[i]
		DB.Model(&models.Operator{}).
			Where("id = ?", operator.ID).
			Updates(operator)

		DB.FirstOrCreate(operator)
	}

}
