package dto

import (
	"encoding/xml"
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
)

type ChinaMobileRsaDecodeRequest struct {
	// EncodeStr 加密字符串
	EncodeStr string `json:"encodeStr" binding:"required"`
}

// ChinaMobileRsaDecodeResponse 响应
type ChinaMobileRsaDecodeResponse struct {
	http.ResponseData
	Data ChinaMobileRsaDecodeResponseData `json:"data"`
}

type ChinaMobileRsaDecodeResponseData struct {
	XMLName    xml.Name `xml:"applyKey" json:"-"`
	Text       string   `xml:",chardata" json:"-"`
	ResponseID string   `xml:"responseID" json:"responseId"`
	ResultCode string   `xml:"resultCode" json:"resultCode"`
	CodeDesc   string   `xml:"codeDesc" json:"codeDesc"`
	ClientCode string   `xml:"clientCode" json:"clientCode"`
	Key        string   `xml:"key" json:"key"`
	Seq        string   `xml:"seq" json:"seq"`
	CreateDate string   `xml:"createDate" json:"createDate"`
	ExpiryDate string   `xml:"expiryDate" json:"expiryDate"`
}
