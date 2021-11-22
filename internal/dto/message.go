package dto

type Message struct {
	Attributes Attributes `json:"attributes"`
	Data       Data       `json:"data"`
}

type Attributes struct {
	Capability string `json:"capability"`
	Channel    string `json:"channel"`
	Commerce   string `json:"commerce"`
	Country    string `json:"country"`
	Datetime   string `json:"datetime"`
	Domain     string `json:"domain"`
	EntityID   string `json:"entityId"`
	EntityType string `json:"entityType"`
	EventID    string `json:"eventId"`
	EventType  string `json:"eventType"`
	MIMEType   string `json:"mimeType"`
	Timestamp  string `json:"timestamp"`
	Version    string `json:"version"`
}

type Data struct {
	Sku           string  `json:"sku"`
	InternalID    string  `json:"internalID"`
	OriginalPrice string  `json:"originalPrice"`
	NormalPrice   string  `json:"normalPrice"`
	ProductCost   string  `json:"productCost"`
	CurrencyCode  string  `json:"currencyCode"`
	Stores        []Store `json:"stores"`
}

type Store struct {
	StoreID          string `json:"storeId"`
	OfferPrice       string `json:"offerPrice"`
	EmployeeDiscount string `json:"employeeDiscount"`
	EmployeePrice    string `json:"employeePrice"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
}
