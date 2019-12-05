package structs

type ReceiveNumber struct {
	CallId string `json:"call_id"`
	Number string `json:"number"`
	Type   int    `json:"type"`
}
