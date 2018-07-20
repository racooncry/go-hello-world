package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const fileName = "E:\\java\\project_相关文件\\海外\\3000备孕用户人群分析2.xlsx"

func main() {

	//readXlsx2()
	readXlsx3()
}

/**
读文件
 */
func readXlsx() {
	excelFileName := fileName
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		checkErr(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}

/**
                             fmt.Printf(",")
							fmt.Printf("%s\n", text)
                            怀孕时间 pregnant_time i==1
 */
func readXlsx2() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ezhome?charset=utf8")
	checkErr(err)

	excelFileName := fileName
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		checkErr(err)
	}

	for i, sheet := range xlFile.Sheets {
		if (i == 1) {
			for j, row := range sheet.Rows {
				if (j > 0) {
					for k, cell := range row.Cells {
						text := cell.String()
						text0 := row.Cells[0].String()
						length := len(row.Cells)
						if (k == 0) {
							fmt.Printf("%s", text)
						}
						if (k == length-1) {
							cell = row.AddCell()
							var pregnant_time string
							//查询数据
							err := db.QueryRow("SELECT pregnant_time FROM app_user where user_id=?", text0).Scan(&pregnant_time)
							checkErr(err)
							fmt.Println("pregnant_time:" + pregnant_time)
							cell.SetString(pregnant_time)
							err = xlFile.Save(fileName)
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}
		}
	}

}



/**
                             fmt.Printf(",")
							fmt.Printf("%s\n", text)
                            怀孕时间 hcgpos>0 i==2
 */
func readXlsx3() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ezhome?charset=utf8")
	checkErr(err)

	excelFileName := fileName
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		checkErr(err)
	}

	for i, sheet := range xlFile.Sheets {
		if (i == 2) {
			for j, row := range sheet.Rows {
				if (j > 0) {
					for k, cell := range row.Cells {
						text := cell.String()
						text0 := row.Cells[0].String()
						length := len(row.Cells)
						if (k == 0) {
							fmt.Printf("%s", text)
						}
						if (k == length-1) {
							cell = row.AddCell()
							var test_date string
							//查询数据
							err := db.QueryRow("SELECT test_date FROM test_strip where user_id=?", text0).Scan(&test_date)
							checkErr(err)
							fmt.Println("test_date:" + test_date)
							cell.SetString(test_date)
							err = xlFile.Save(fileName)
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}
		}
	}

}








/**
打开文件，并向里面写内容
 */
func writeXlsx() {
	file, err := xlsx.OpenFile(fileName)
	if err != nil {
		panic(err)
	}
	first := file.Sheets[0]
	row := first.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "1"
	cell = row.AddCell()
	cell.Value = "张三"
	cell = row.AddCell()
	cell.Value = "男"
	cell = row.AddCell()
	cell.Value = "18"

	err = file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
