package database

import(

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"fmt"
	"os"
)
var DB *gorm.DB
func GormConnect() {
		//DBMS     := "mysql"
		user := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := "3306"
		database_name := os.Getenv("MYSQL_DATABASE")

		CONNECT := user+":"+password+"@tcp("+host+":"+port+")/"+database_name+"?charset=utf8mb4"
		DB,_ = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

		//if err != nil {
		//	fmt.Println(err)
		//}
}