package vodacom

import "encoding/xml"

type MpesaBroker struct {
	XMLName        xml.Name       `xml:"mpesaBroker"`
	Xmlns          string         `xml:"xmlns,attr"`
	Version        string         `xml:"version,attr"`
	RequestElement RequestElement `xml:"request"`
}

type RequestElement struct {
	ServiceProvider ServiceProvider `xml:"serviceProvider"`
	Transaction     Transaction     `xml:"transaction"`
}

type ServiceProvider struct {
	SpID       string `xml:"spId"`
	SpPassword string `xml:"spPassword"`
	Timestamp  string `xml:"timestamp"`
}

type Transaction struct {
	Amount                 float64 `xml:"amount"`
	CommandID              string  `xml:"comandId"`
	Initiator              string  `xml:"initiator"`
	OriginatorConversation string  `xml:"originatorConversationID"`
	Recipient              string  `xml:"recipient"`
	MpesaReceipt           string  `xml:"mpesaReceipt"`
	TransactionDate        string  `xml:"transactionDate"`
	AccountReference       string  `xml:"accountReference"`
	TransactionID          string  `xml:"transactionID"`
	ConversationID         string  `xml:"conversationID"`
}
