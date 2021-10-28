package configration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//var DB *sql.DB
func OpenConnection()*sql.DB{
	driver:="mysql"
	congif:="root:00@/newResturant"
	db,err:=sql.Open(driver,congif)
	if err!=nil {
		fmt.Print("error while establishing connection ",err)
	}
	return db
}