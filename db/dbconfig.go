package db

import (
	// "com/josh/asset/service"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "mysql#3306"
	dbname   = "AssetManagement"
)

var db *gorm.DB

func InitDatabase() {
	// init database
	MysqlConnection()
}

func MysqlConnection() {

	info := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	// database, err := sql.Open("mysql", info)
	database, err := gorm.Open("mysql", info)

	if err != nil {
		panic(err.Error())
	}
	db = database
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	panic(err)
	// }
	// close db when not in use
	// defer db.Close()

	// Migrate the schema
	// db.AutoMigrate(
	// 	&service.User{})

	fmt.Println("Successfully connected!", db)

}

func GetDB() *gorm.DB {
	return db
}
