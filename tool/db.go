package tool

import(
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Pk string
	Columns map[string]interface{} 
}

func NewDb() *Db {
	obj := new(Db)
	return obj
}

func (this *Db) Conn() *sql.DB {
	config := NewConfig()
	host := config.Cfg("host")
	port := config.Cfg("port")
	user := config.Cfg("user")
	password := config.Cfg("password")
	dbname := config.Cfg("dbname")
	conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	db, err := sql.Open("mysql", conn)
	if err != nil{
		fmt.Println(err)
	}
	
	return db
	// defer db.Close()
}

func (this *Db) SetColumn(tableName string) {
	db := this.Conn()
	defer db.Close()
	config := NewConfig()
	prefix := config.Cfg("perfix")
	tableName = prefix + tableName
	tableName = strings.ToLower(tableName)
	rows, _ := db.Query("select COLUMN_NAME,DATA_TYPE,COLUMN_KEY from information_schema.columns where table_name=?", tableName)
	defer rows.Close()

	columns := make(map[string]interface{}, 0)

	for rows.Next() {
		var name, data_type, column_key string

		if err := rows.Scan(&name, &data_type, &column_key); err != nil {
			fmt.Println(err)
		}

		if column_key == "PRI" {
			this.Pk = name
		}

		columns[name] = data_type
	}
	this.Columns = columns

}
