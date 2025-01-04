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
	//var ip, pwd string
	//fmt.Println("输入ip和密码")
	//_, err = fmt.Scanln(&ip)
	//if err != nil {
	//	return
	//}
	//_, err = fmt.Scanln(&pwd)
	//if err != nil {
	//	return
	//}
	// 链接 postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		global.PostgresIP /*ip*/, global.PostgresUsr, global.PostgresPwd, /*pwd*/
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
		&models.User{},     // 账户
		&models.Cart{},     //购物车
		&models.Product{},  //商品
		&models.Category{}, //商品类别
		&models.Order{},    // 订单
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功，模型迁移成功")

}
