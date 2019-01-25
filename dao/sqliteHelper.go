package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func RunToList(strSQL string, iPageIndex int, iPageSize int, iPageCount *int) ([]map[string]string, error) {
	db, _ := getdb()
	strTotalSQL := "select count(1) as RowCount from ( " + strSQL + " )x"

	var iTotalRows int
	err := db.QueryRow(strTotalSQL).Scan(&iTotalRows)

	if err != nil {
		log.Fatal(err)
	}

	if iTotalRows%iPageSize == 0 {
		*iPageCount = iTotalRows / iPageSize
	} else {
		*iPageCount = iTotalRows/iPageSize + 1
	}

	strSQL = strSQL + " limit " + strconv.Itoa(iPageSize) + " offset " + strconv.Itoa((iPageIndex-1)*iPageSize)

	rows, err := db.Query(strSQL)

	fmt.Println(strSQL)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
		fmt.Println(columns)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var rv []map[string]string

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		rowMap := make(map[string]string)
		var value string
		for i, col := range values {
			if col != nil {
				value = string(col)
				rowMap[columns[i]] = value
			}
		}
		rv = append(rv, rowMap)
	}

	return rv, err
}
