package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mark3/global"
	model "mark3/repository/db/model/rule"
	"mark3/service/rule_action_log"
	"strconv"
	"strings"
	//"strconv"
)

var cn = cron.New(cron.WithSeconds())
var IsStarted = 0
var CronMap map[string]CronItem

type CronItem struct {
	Id   int
	Spec string
}

func Add(mainId, spec string, id int) {
	CronMap[mainId] = CronItem{
		Id:   id,
		Spec: spec,
	}
}

func Delete(mainId string) {
	delete(CronMap, mainId)
}

func Get(mainId string) (CronItem, bool) {
	_map, ok := CronMap[mainId]
	return _map, ok
}

func StopAll() {
	fmt.Println("stop all...")
	IsStarted = 0
	cn.Stop()
}
func StartAll() {
	fmt.Println("start all...")
	IsStarted = 1
	cn.Start()
}

func InitScheduleAll(isStop int) {
	fmt.Println("init all...")
	CronMap = map[string]CronItem{}
	var ruleList []model.RuleModel
	global.GVA_DB.Model(model.RuleModel{}).Select("*").Where("rule_type='schedule' and is_enable='yes' and is_deleted='no' and ws_id=46 and tmpl_id=174 ").Find(&ruleList)
	for _, rule := range ruleList {
		AddTaskCron(rule)
	}
	if isStop != 1 {
		IsStarted = 1
		cn.Start()
	}
}

func AddTaskCron(rule model.RuleModel) {
	if rule.Id < 1 || rule.TriggerTime == "" || rule.FieldKey == "" {
		return
	}
	if rule.RuleType == "schedule" && rule.IsDeleted == "no" && rule.IsEnable == "yes" {
		times := strings.Split(rule.TriggerTime, ":")
		spec := fmt.Sprintf("0 %s %s * * ?", times[1], times[0])
		id, _ := cn.AddFunc(spec, func() {
			ExecuteTaskCron(rule)
		})
		Add(strconv.Itoa(rule.Id), spec, int(id))
	}
}

func DeleteTaskCron(rule model.RuleModel) {
	if rule.Id > 0 {
		cmap, ok := Get(strconv.Itoa(rule.Id))
		if ok {
			cn.Remove(cron.EntryID(cmap.Id))
			Delete(strconv.Itoa(rule.Id))
		}
	}
}

func UpdateTaskCron(rule model.RuleModel) {
	DeleteTaskCron(rule)
	AddTaskCron(rule)
}

func ExecuteTaskCron(rule model.RuleModel) {
	// 查询mongo数据
	documents, err := rule_action_log.GetDocuments(rule)
	if err != nil {
		return
	}
	// 遍历数据，执行动作
	for _, document := range documents {
		if id, ok := document["_id"]; !ok {
			continue
		} else {
			objID := id.(primitive.ObjectID)
			issueId := objID.Hex()
			rule_action_log.AddRuleActionLog(rule, issueId, document)
		}
	}
}
