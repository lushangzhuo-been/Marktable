package tmpl

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/message"
	"mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/repository/db/model/ws"
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gopkg.in/gomail.v2"
)

func reassignNotice(wsId int, tmplId int, issueId string, userid int, toUserList []int, title string) {
	var ws ws.WsModel
	if err := global.GVA_DB.Where("id=?", wsId).First(&ws).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	var tmpl tmpl.TmplModel
	if err := global.GVA_DB.Where("id=?", tmplId).First(&tmpl).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var user user.UserModel
	if err := global.GVA_DB.Where("id=?", userid).First(&user).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	logoURL := global.GVA_CONFIG.URL.Logo
	avatarURL := fmt.Sprintf(global.GVA_CONFIG.URL.Avatar, user.Avatar)
	issueURL := fmt.Sprintf(global.GVA_CONFIG.URL.Issue, wsId, tmplId, issueId)

	//向消息中心发送消息
	var content enum.TemplateForReassign
	content.IssueTitle = title
	content.IssueURL = issueURL
	content.WsName = ws.Name
	content.TmplName = tmpl.Name
	c, _ := json.Marshal(content)

	var messageList []interface{}
	for _, toUserid := range toUserList {
		message := bson.M{}
		message["userid"] = toUserid
		message["content"] = string(c)
		message["origin_userid"] = userid
		message["mode"] = enum.ModeReassign
		message["read_sign"] = "no"
		message["created_at"] = common.GetCurrentTime()
		messageList = append(messageList, message)
	}
	collection := global.GVA_MONGO.Database("mark3").Collection("user_message")
	_, err := collection.InsertMany(context.TODO(), messageList)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	//发送邮件
	t, err := template.ParseFiles("./views/email.html")
	if err != nil {
		global.GVA_LOG.Error("create tpl failed, err :", err.Error())
		return
	}
	type Reassign struct {
		WsId      int
		TmplId    int
		IssueId   string
		WsName    string
		TmplName  string
		Creator   string
		Title     string
		LogoURL   string
		AvatarURL string
		IssueURL  string
	}
	var emailTpl bytes.Buffer
	if err = t.Execute(&emailTpl, &Reassign{
		WsId:      wsId,
		TmplId:    tmplId,
		IssueId:   issueId,
		WsName:    ws.Name,
		TmplName:  tmpl.Name,
		Creator:   user.FullName,
		Title:     title,
		LogoURL:   logoURL,
		AvatarURL: avatarURL,
		IssueURL:  issueURL,
	}); err != nil {
		global.GVA_LOG.Error("execute tpl failed, err :", err.Error())
		return
	}

	subject := "[项目管理]任务通知，您有一条新任务了"
	body := emailTpl.String()
	var to []string

	userInfoList := GetUserList(toUserList)
	for _, userInfo := range userInfoList {
		email := userInfo.Email
		if email == "" {
			continue
		}
		to = append(to, email)
	}
	if err := SendEmail(to, nil, nil, subject, body, ""); err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}

func AtPersonNotice(userid int, wsId int, tmplId int, progress bson.M) {
	progressContent := progress["content"].(string)
	issueId := progress["issue_id"].(string)

	pattern, _ := regexp.Compile(`<span class="mention-at" id="(.*?)" contenteditable="false">`)
	result := pattern.FindAllStringSubmatch(progressContent, -1)
	var toUserList []int
	for _, line := range result {
		toUserid, err := strconv.Atoi(line[1])
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		toUserList = append(toUserList, toUserid)
	}
	if len(toUserList) == 0 {
		return
	}

	var ws ws.WsModel
	if err := global.GVA_DB.Where("id=?", wsId).First(&ws).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	var tmpl tmpl.TmplModel
	if err := global.GVA_DB.Where("id=?", tmplId).First(&tmpl).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		return
	}
	filter := bson.M{"_id": objectID}

	var issueDocument = bson.M{}
	err = collection.FindOne(context.TODO(), filter).Decode(&issueDocument)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	//向消息中心发送消息
	issueURL := fmt.Sprintf(global.GVA_CONFIG.URL.Issue, wsId, tmplId, issueId)

	var template enum.TemplateForAtPerson
	template.ProgressContent = progressContent
	template.IssueTitle = issueDocument["title"].(string)
	template.IssueURL = issueURL
	template.WsName = ws.Name
	template.TmplName = tmpl.Name
	t, _ := json.Marshal(template)

	var messageList []interface{}
	for _, toUserid := range toUserList {
		message := bson.M{}
		message["userid"] = toUserid
		message["content"] = string(t)
		message["origin_userid"] = userid
		message["mode"] = enum.ModeAtPerson
		message["read_sign"] = enum.ReadSignNo
		message["created_at"] = common.GetCurrentTime()
		messageList = append(messageList, message)
	}
	collection = global.GVA_MONGO.Database("mark3").Collection("user_message")
	_, err = collection.InsertMany(context.TODO(), messageList)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

}

func SendEmail(to []string, cc []string, bcc []string, subject string, body string, attach string) error {
	if global.GVA_CONFIG.Smtp.IsOpen != "yes" {
		return fmt.Errorf("邮件服务器未开启")
	}

	host := global.GVA_CONFIG.Smtp.Host
	port := global.GVA_CONFIG.Smtp.Port
	username := global.GVA_CONFIG.Smtp.Username
	password := global.GVA_CONFIG.Smtp.Password
	from := global.GVA_CONFIG.Smtp.From

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if len(cc) > 0 {
		//抄送
		m.SetHeader("Cc", cc...)
	}
	if len(bcc) > 0 {
		// 密送
		m.SetHeader("Bcc", bcc...)
	}
	if len(attach) > 0 {
		//添加附件
		m.Attach(attach)
	}

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		global.GVA_LOG.Error(err.Error())
		return err
	}
	return nil
}
