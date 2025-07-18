package tmpl

import (
	"fmt"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/service/tmpl"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func exportExcel(ctx *gin.Context, tmpl model.TmplModel, documents []bson.M, tmplCaller *tmpl.Caller) {
	fields, _ := tmplCaller.GetFields(tmpl.Id)
	var titleList []string
	var keyList []string
	for _, field := range fields {
		titleList = append(titleList, field.Name)
		keyList = append(keyList, field.FieldKey)
	}
	var data [][]string
	for _, document := range documents {
		documentFormat := formatDocument(tmpl.Id, document, fields, tmplCaller)
		var tmp []string
		for _, key := range keyList {
			tmp = append(tmp, documentFormat[key])
		}
		data = append(data, tmp)
	}
	f := excelize.NewFile()
	sheet, _ := f.NewSheet("Sheet1")
	f.SetActiveSheet(sheet)
	f.SetSheetRow("sheet1", "A1", &titleList)
	row := 1
	for _, v := range data {
		row++
		f.SetSheetRow("sheet1", fmt.Sprintf("A%d", row), &v)
	}

	fileName := fmt.Sprintf("%s_%s.xlsx", tmpl.Name, time.Now().Format("20060102150405"))

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", url.PathEscape(fileName)))
	ctx.Header("filename", url.PathEscape(fileName))
	ctx.Header("Access-Control-Expose-Headers", "filename")
	f.Write(ctx.Writer)
}

func formatDocument(tmplId int, document bson.M, fields []model.FieldModel, tmplCaller *tmpl.Caller) (documentFormat map[string]string) {
	var fieldsMap = make(map[string]model.FieldModel)
	for _, field := range fields {
		fieldsMap[field.FieldKey] = field
	}
	documentFormat = make(map[string]string)
	for key, value := range document {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}

		switch field.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom:
			valueAssert := value.(primitive.A)
			var strArr []string
			for _, v := range valueAssert {
				vStr, ok := v.(string)
				if !ok {
					continue
				}
				strArr = append(strArr, vStr)
			}
			documentFormat[key] = strings.Join(strArr, ",")
		case enum.FieldNumberCom, enum.FieldProgressCom:
			vInt32, ok := value.(int32)
			var str string
			if ok {
				str = fmt.Sprintf("%d", vInt32)
			}
			documentFormat[key] = str
		case enum.FieldPersonCom:
			valueAssert := value.(primitive.A)
			var intArr []int
			for _, v := range valueAssert {
				vInt32, ok := v.(int32)
				if !ok {
					continue
				}
				intArr = append(intArr, int(vInt32))
			}
			users := GetUserList(intArr)
			var fullNameArr []string
			for _, user := range users {
				fullNameArr = append(fullNameArr, user.FullName)
			}
			documentFormat[key] = strings.Join(fullNameArr, ",")
		case enum.FieldLinkCom:
			valueAssert, ok := value.(primitive.M)
			if !ok {
				continue
			}
			var strArr []string
			for k, v := range valueAssert {
				vStr := v.(string)
				if k == "url" && vStr != "" {
					strArr = append(strArr, fmt.Sprintf("网址:%s", vStr))
				}
				if k == "name" && vStr != "" {
					strArr = append(strArr, fmt.Sprintf("名称:%s", vStr))
				}
			}
			documentFormat[key] = strings.Join(strArr, ",")
		case enum.FieldStatusCom:
			statusId := value.(int32)
			var str string
			status, err := tmplCaller.GetStatus(tmplId, int(statusId))
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
			documentFormat[key] = str
		case enum.FieldRelatedCom:
			var str string
			valueAssert := value.(primitive.A)
			for _, issueId := range valueAssert {
				issueIdStr, okr := issueId.(string)
				if okr {
					relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
					if err1 == nil {
						str = fmt.Sprintf("%s", relatedDocument["title"])
						break
					}
				}
			}
			documentFormat[key] = str
		default:
			documentFormat[key] = value.(string)
		}
	}
	return
}
