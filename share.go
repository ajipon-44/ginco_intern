package share

import(
    //"database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux"
		"github.com/jinzhu/gorm"
		//"github.com/jinzhu/gorm/dialects/mysql"
		"log"
)

func ConnectDb() *gorm.DB {
	dbconf := "root:Yanakei727@tcp(127.0.0.1:3306)/ginco_intern"
	db, err_db := gorm.Open("mysql", dbconf)
	if err_db != nil {
		panic(err_db.Error())
	}

	return db
}
