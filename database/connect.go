package database

import(

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"os"
)
var DB *gorm.DB
func GormConnect() {
	var err error
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := "3306"
	database_name := os.Getenv("MYSQL_DATABASE")

	CONNECT := user+":"+password+"@tcp("+host+":"+port+")/"+database_name+"?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	sqlDB, err := DB.DB()
	err = sqlDB.Ping()

	if err != nil {
		fmt.Println("エラー")
		fmt.Println(err)
		return
	} else {
		fmt.Println("データベース接続成功")
	}
}