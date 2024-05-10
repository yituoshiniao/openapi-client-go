package dto

import "github.com/yituoshiniao/gin-api-http/internal/api/http"

type ExampleGetRequest struct {
	//user_id 用户id
	UserID string `form:"user_id"   binding:"required"`
	//query_id
	QueryID string `form:"query_id" binding:"required"`
	//create_time
	CreateTime int32 `form:"create_time" binding:"required"`
}

type ExampleGetResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type ExampleGetOneRequest struct {
	//user_id 用户id
	UserID string `form:"user_id"   binding:"required"`
	//query_id
	QueryID string `form:"query_id" binding:"required"`
	//create_time
	CreateTime int32 `form:"create_time" binding:"required"`
}

type ExampleGetOneResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type ExamplePostRequest struct {
	//CmdId 功能类型枚举:
	//1 AUTH鉴权;
	//2 QA 问答；
	//4 总结SUMMARY;
	//5 KEYWORDS 关键词；
	//6 CLASSIFICATION 文档分类；
	//7 RENAME  重命名；
	//8 mindmap 思维导图；
	CmdId int8 `json:"cmd_id" binding:"required" example:"1" binding:""`
	//Token 鉴权token
	Token string `json:"token" binding:"required" example:"sdfsdsdfsd" binding:""`
	//CVer 客户端协议版本号
	CVer string `json:"c_ver" binding:"required"`
	//AppID 应用id
	AppID string `json:"app_id" binding:"required"`
	//Fid： 文档id
	Fid string `json:"fid" binding:"required"`
	// Lang 语言
	Lang string `json:"lang" binding:"required"`
	// Question 问题
	Question string `json:"question" binding:"required"`
	// show_gpt_2
	ShowGpt2 string `json:"show_gpt_2"`
	// IgnoreFeaturePrompt 是否忽略功能检查  1 是； 0 否；
	IgnoreFeaturePrompt int8 `json:"ignore_feature_prompt"`
}

// ExamplePostResponse 示例返回
type ExamplePostResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type UserPortraitData struct {
	//用户id
	UserId string `json:"user_id"`
	//上次登陆时间
	LastLogin string `json:"last_login"`
	//国家
	Country string `json:"country"`
	//是否为VIP，0/1
	VipInfo int `json:"vip_info"`
}
