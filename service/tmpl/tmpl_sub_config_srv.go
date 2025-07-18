package tmpl

import (
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"
	"strconv"
	"strings"
)

type TmplSubConfigSrv struct{}

func (s *TmplSubConfigSrv) TmplSubConfigCheck(req types.TmplSubConfigCheckReq) (resp interface{}, err error) {
	var limits []model.TmplSubConfigModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=?", req.WsId, req.TmplId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}

	var tmplIds []int
	tmplIds = append(tmplIds, req.TmplId)
	for _, v := range limits {
		tmplIds = append(tmplIds, v.SubTmplId)
	}
	type TmplVo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var tmplList []TmplVo
	global.GVA_DB.Model(&model.TmplModel{}).Where("ws_id=? and id not in ?", req.WsId, tmplIds).Scan(&tmplList)

	return tmplList, nil
}

func (s *TmplSubConfigSrv) TmplSubConfigList(req types.TmplSubConfigListReq) (resp interface{}, err error) {
	type TmplSubConfigVo struct {
		Id          int    `json:"id"`
		WsId        int    `json:"ws_id"`
		TmplId      int    `json:"tmpl_id"`
		SubTmplId   int    `json:"sub_tmpl_id"`
		SubTmplName string `json:"sub_tmpl_name"`
	}
	var tmplSubConfigList []TmplSubConfigVo
	sql := "select a.id, a.ws_id, a.tmpl_id, a.sub_tmpl_id, b.name as sub_tmpl_name from tmpl_sub_config AS a LEFT JOIN tmpl AS b ON a.sub_tmpl_id = b.id where a.ws_id=? and a.tmpl_id=?"
	global.GVA_DB.Raw(sql, req.WsId, req.TmplId).Scan(&tmplSubConfigList)

	// 排序
	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, enum.ScreenCoordinateModuleSubTmpl).First(&coordinate)

	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var dataMap = make(map[string]TmplSubConfigVo)
		for _, vo := range tmplSubConfigList {
			dataMap[strconv.Itoa(vo.SubTmplId)] = vo
		}

		var sort []TmplSubConfigVo
		for _, key := range coordinateSlice {
			if data, ok := dataMap[key]; ok {
				sort = append(sort, data)
				delete(dataMap, key)
			}
		}

		for _, vo := range dataMap {
			sort = append(sort, vo)
		}
		return sort, nil
	}

	return tmplSubConfigList, nil
}

func (s *TmplSubConfigSrv) TmplSubConfigCreate(req types.TmplSubConfigCreateReq) (resp interface{}, err error) {
	var limits []model.TmplSubConfigModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and sub_tmpl_id=?", req.WsId, req.TmplId, req.SubTmplId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	if len(limits) > 0 {
		return nil, errors.New("不能重复添加")
	}

	if err = global.GVA_DB.Create(&model.TmplSubConfigModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		SubTmplId: req.SubTmplId,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *TmplSubConfigSrv) TmplSubConfigDelete(req types.TmplSubConfigDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where(
		"ws_id=? and tmpl_id=? and sub_tmpl_id=?", req.WsId, req.TmplId, req.SubTmplId).Delete(
		&model.TmplSubConfigModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
