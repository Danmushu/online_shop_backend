package postgres

import (
	"fmt"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/global"
	"project/models"
)

const port = 5432

var DB *gorm.DB

// todo
func init() {
	var err error

	// 链接 postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		global.PostgresIP, global.PostgresUsr, global.PostgresPwd,
		global.PostgresDB, port)
	DB, err = gorm.Open(driver.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}, //自动迁移时，禁用外键约束,不禁用
	)

	if err != nil {
		// 出错
		fmt.Println("数据库连接失败")
		fmt.Println(err)
		return
	}

	//迁移模型
	err = DB.AutoMigrate(
		&models.User{}, // 账户
		//&models.Operator{}, // 管理员
		&models.Cart{},     //购物车
		&models.Product{},  //商品
		&models.Category{}, //商品类别
		&models.Order{},    // 订单
	)

}
