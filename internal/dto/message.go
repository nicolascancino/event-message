package dto

type Message struct {
	Attributes map[string]string `json:"attributes"`
	Data       interface{}       `json:"data"`
}
