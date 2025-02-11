package model

type HashRequest struct {
	Text      string `json:"text"`
	Algorithm string `json:"algorithm"`
}

type HashResponse struct {
	Original  string `json:"original"`
	Hashed    string `json:"hashed"`
	Algorithm string `json:"algorithm"`
}
