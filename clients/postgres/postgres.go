package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"project/global"
)

const port = 5432

var DB *gorm.DB

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
		&models.ProChange{}, // 变化
		&models.Machine{},   // 账户
		&models.Code{},      // 激活码
		&models.Operator{},  // 管理员
		&models.Order{},     // 订单
		&models.Gift{},      // 全体赠送
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
