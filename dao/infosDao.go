package dao

import (
	"log"
	"strconv"

	mdl "51hk_beego/models"
)

//GetAll  to get all infomations
func GetAll(id int, iCatgId int, strKeyword string, iPageIndex int, iPageSize int, iPageCount *int) (mdl.InfoList, error) {

	strSQL := "select a.*,b.name as CatgName  from infos a left join category b on a.catgid=b.id where 1=1 "

	if id != -1 {
		strSQL = strSQL + " and a.Id=" + strconv.Itoa(id)
	}
	if iCatgId != -1 {
		strSQL = strSQL + " and a.CatgId=" + strconv.Itoa(iCatgId)
	}

	if strKeyword != "" {
		strSQL = strSQL + " and (a.content like ''" + strKeyword + "'' or a.title like ''%" + strKeyword + "%''"
	}

	strSQL = strSQL + (" order by a.updatetime desc")
	log.Println(strSQL)

	infosMap, err := RunToList(strSQL, iPageIndex, iPageSize, iPageCount)

	var rv mdl.InfoList
	rv.BuildModelList(infosMap)

	return rv, err
}
