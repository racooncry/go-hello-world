package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"

)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "123456"         //密码
	dbname     = "go_demo"        //表名
)

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT person SET name=?,age=?")
	checkErr(err)

	res, err := stmt.Exec("码农", 24)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)



	//更新数据
	stmt, err = db.Prepare("update person set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("码农三代", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM person")
	checkErr(err)
	fmt.Println("id    name    age")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		checkErr(err)
		fmt.Print(id,"    ")
		fmt.Print(name,"    ")
		fmt.Println(age)

	}




	//删除数据
	stmt, err = db.Prepare("delete from person where name=?")
	checkErr(err)

	res, err = stmt.Exec("码农")
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()


}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}