package models

import (
	"strconv"
)

//InfoModel is info struct
type InfoModel struct {
	Id           int    `db:"id"`
	Title        string `db:"Title"`
	ShortDesc    string `db:"ShortDesc"`
	CatgId       int    `db:"CatgId"`
	CatgName     string `db:"CatgName"`
	Content      string `db:"Content"`
	CreateUserId int    `db:"CreateUserId"`
	CreateTime   string `db:"CreateTime"`
	ReadCount    int    `db:"ReadCount"`
	UpdateTime   string `db:"UpdateTime"`
}

type InfoList struct {
	Data []InfoModel
}

type Convert interface {
	BuildModel()
	BuildModelList()
}

func (info *InfoModel) BuildModel(srcData map[string]string) {
	info.Id, _ = strconv.Atoi(srcData["Id"])
	info.Title = srcData["Title"]
	info.ShortDesc = srcData["ShortDesc"]
	info.CatgId, _ = strconv.Atoi(srcData["CatgId"])
	info.CatgName = srcData["CatgName"]
	info.Content = srcData["Content"]
	info.CreateUserId, _ = strconv.Atoi(srcData["CreateUserId"])
	info.CreateTime = srcData["CreateTime"]
	info.ReadCount, _ = strconv.Atoi(srcData["ReadCount"])
	info.UpdateTime = srcData["UpdateTime"]
}

func (infoLst *InfoList) BuildModelList(srcData []map[string]string) {

	for _, singData := range srcData {
		var infoMode InfoModel
		infoMode.Id, _ = strconv.Atoi(singData["Id"])
		infoMode.Title = singData["Title"]
		infoMode.ShortDesc = singData["ShortDesc"]
		infoMode.CatgId, _ = strconv.Atoi(singData["CatgId"])
		infoMode.CatgName = singData["CatgName"]
		infoMode.Content = singData["Content"]
		infoMode.CreateUserId, _ = strconv.Atoi(singData["CreateUserId"])
		infoMode.CreateTime = singData["CreateTime"]
		infoMode.ReadCount, _ = strconv.Atoi(singData["ReadCount"])
		infoMode.UpdateTime = singData["UpdateTime"]

		infoLst.Data = append(infoLst.Data, infoMode)
	}
}
