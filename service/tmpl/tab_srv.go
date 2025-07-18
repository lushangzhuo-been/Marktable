package tmpl

import (
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/types/common"
	types "mark3/types/tmpl"
	"strings"
)

type TabSrv struct{}

func (s *TabSrv) TabConfig(req types.TmplTabConfigReq) (resp interface{}, err error) {
	// 已选中
	var limits []model.TmplTabModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=?", req.WsId, req.TmplId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	var checkedTab = make(map[string]string)
	for _, v := range limits {
		checkedTab[v.Tab] = v.Tab
	}
	// 流程详情 默认已添加
	if _, ok := checkedTab[types.TabTmplDetail]; !ok {
		checkedTab[types.TabTmplDetail] = types.TabTmplDetail
	}

	type TmplTabNotCheckVo struct {
		common.BaseConfig
		Checked bool `json:"checked"`
	}

	var res []TmplTabNotCheckVo
	for _, config := range types.TabConfig {
		if _, ok := checkedTab[config.Value]; ok {
			res = append(res, TmplTabNotCheckVo{BaseConfig: config, Checked: true})
		} else {
			res = append(res, TmplTabNotCheckVo{BaseConfig: config, Checked: false})
		}
	}

	return res, nil
}

func (s *TabSrv) TmplTabList(req types.TmplTabListReq) (resp interface{}, err error) {
	type TmplTabWithTabVo struct {
		model.TmplTabModel
		TabCn interface{} `json:"tab_cn"`
	}
	var limits []model.TmplTabModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=?", req.WsId, req.TmplId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}

	var res []TmplTabWithTabVo
	isTmplDetail := false
	for _, v := range limits {
		if v.Tab == types.TabTmplDetail {
			isTmplDetail = true
		}
		res = append(res, TmplTabWithTabVo{TmplTabModel: v, TabCn: types.TabMap[v.Tab]})
	}
	if !isTmplDetail {
		var tmplTabDetail = model.TmplTabModel{
			WsId:   req.WsId,
			TmplId: req.TmplId,
			Tab:    types.TabTmplDetail,
		}
		global.GVA_DB.Create(&tmplTabDetail)
		res = append(res, TmplTabWithTabVo{TmplTabModel: tmplTabDetail, TabCn: types.TabMap[tmplTabDetail.Tab]})
	}

	// 排序
	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, enum.ScreenCoordinateModuleTab).First(&coordinate)
	var coordinateSlice []string
	if coordinate.Id != 0 {
		coordinateSlice = strings.Split(coordinate.Coordinate, ",")
	} else {
		for _, config := range types.TabConfig {
			coordinateSlice = append(coordinateSlice, config.Value)
		}
	}

	var dataMap = make(map[string]TmplTabWithTabVo)
	for _, vo := range res {
		dataMap[vo.Tab] = vo
	}

	var sort []TmplTabWithTabVo
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

func (s *TabSrv) TmplTabCreate(req types.TmplTabCreateReq) (resp interface{}, err error) {
	var limits []model.TmplTabModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and tab=?", req.WsId, req.TmplId, req.Tab).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	if len(limits) > 0 {
		return nil, errors.New("不能重复添加")
	}

	if err = global.GVA_DB.Create(&model.TmplTabModel{
		WsId:   req.WsId,
		TmplId: req.TmplId,
		Tab:    req.Tab,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *TabSrv) TmplTabDelete(req types.TmplTabDeleteReq) (resp interface{}, err error) {
	if req.Tab == types.TabTmplDetail {
		return nil, errors.New("流程详情禁止删除")
	}

	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and tab=?", req.WsId, req.TmplId, req.Tab).Delete(&model.TmplTabModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
