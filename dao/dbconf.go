package dao

import (
	"database/sql"
	"os"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func getdb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "E:\\jkl\\liode\\goProjects\\src\\51hk_beego\\app_data\\jzh")
	return db, err
}

func getDbByName(dbname string) (*sql.DB, error) {
	dbPath := getCurrentPath() + "\\" + dbname
	db, err := sql.Open("sqlite3", dbPath)
	return db, err
}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
