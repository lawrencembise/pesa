package vodacom

import "encoding/xml"

type MpesaBrokerResponseConfirming struct {
	XMLName         xml.Name                  `xml:"mpesaBroker"`
	Xmlns           string                    `xml:"xmlns,attr"`
	Version         string                    `xml:"version,attr"`
	ResponseElement ResponseElementConfirming `xml:"response"`
}

type ResponseElementConfirming struct {
	ConversationID         string `xml:"conversationID"`
	OriginatorConversation string `xml:"originatorConversationID"`
	ResponseCode           int    `xml:"responseCode"`
	ResponseDesc           string `xml:"responseDesc"`
	ServiceStatus          string `xml:"serviceStatus"`
	TransactionID          string `xml:"transactionID"`
}
