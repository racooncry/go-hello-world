package main


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_demo"
)

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	fmt.Println("begin connect mysql....")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		panic(err)
		return
	}
	fmt.Println("connnect success!")
}

func InsertTx() {
	//开启事务
	tx, err := DB.Begin()
	checkErr(err)

	//准备sql语句
	//插入数据

	stmt, err := tx.Prepare("INSERT person SET name=?,age=?")
	fmt.Println("check insert sql2")
	checkErr(err)
	res, err := stmt.Exec("码农", 24)
	checkErr(err)

	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
}

func Insert() {
	//准备sql语句
	//插入数据
	stmt, err := DB.Prepare("INSERT person SET name=?,age=?")
	checkErr(err)
	res, err := stmt.Exec("码农", 24)
	checkErr(err)
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
}

func main() {

	InitDB()
	Insert()


}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}