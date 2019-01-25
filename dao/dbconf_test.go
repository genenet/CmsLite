package dao

import (
	"fmt"
	"testing"
)

func Test_getDbByName(t *testing.T) {
	db, _ := getDbByName("litemain.db")
	fmt.Print(db)
}
