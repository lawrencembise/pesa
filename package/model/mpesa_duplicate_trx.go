package model

import "encoding/xml"

type MpesaBrokerResult struct {
	XMLName       xml.Name      `xml:"mpesaBroker"`
	Xmlns         string        `xml:"xmlns,attr"`
	Version       string        `xml:"version,attr"`
	ResultElement ResultElement `xml:"result"`
}

type ResultElement struct {
	ServiceProvider ServiceProviderResult `xml:"serviceProvider"`
	Transaction     TransactionResult     `xml:"transaction"`
}

type ServiceProviderResult struct {
	SpID       string `xml:"spId"`
	SpPassword string `xml:"spPassword"`
	Timestamp  string `xml:"timestamp"`
}

type TransactionResult struct {
	ResultType             string `xml:"resultType"`
	ResultCode             int    `xml:"resultCode"`
	ResultDesc             string `xml:"resultDesc"`
	ServiceReceipt         string `xml:"serviceReceipt"`
	ServiceDate            string `xml:"serviceDate"`
	OriginatorConversation string `xml:"originatorConversationID"`
	ConversationID         string `xml:"conversationID"`
	TransactionID          string `xml:"transactionID"`
	Initiator              string `xml:"initiator"`
	InitiatorPassword      string `xml:"initiatorPassword"`
}
