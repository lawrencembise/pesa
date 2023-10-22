package model

import "encoding/xml"

type MpesaBrokerResponse struct {
	XMLName         xml.Name        `xml:"mpesaBroker"`
	Xmlns           string          `xml:"xmlns,attr"`
	Version         string          `xml:"version,attr"`
	ResponseElement ResponseElement `xml:"response"`
}

type ResponseElement struct {
	ConversationID         string `xml:"conversationID"`
	OriginatorConversation string `xml:"originatorConversationID"`
	TransactionID          string `xml:"transactionID"`
	ResponseCode           int    `xml:"responseCode"`
	ResponseDesc           string `xml:"responseDesc"`
	ServiceStatus          string `xml:"serviceStatus"`
}
