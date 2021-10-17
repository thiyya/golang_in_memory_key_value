package models

type Item struct {
	// Value
	Value string
	// Time to live duration
	TTL int64
}

// Key Value Item
//
// swagger:model
type KeyValueItem struct {
	// Key
	Key string
	// Value
	Value string
	// Time to live duration
	Ttl int64
}

// ValueResponse Key'e ait değer
// swagger:response
type ValueResponse struct {
	// Value
	Value string `json: "value"`
}

// ErrorResponse400 Hatalı istek
// swagger:response
type ErrorResponse400 struct {
	// Hata Mesaji
	Message string `json: "message"`
}

// GenericErrorResponse Diğer hatalar
// swagger:response
type GenericErrorResponse struct {
	// Hata Kodu
	Code string `json: "code"`
	// Hata Mesaji
	Message string `json: "message"`
}

// KeyValueItemResponse Kaydedilen Key Value Item
// swagger:response
type KeyValueItemResponse struct {
	// Key Value Item
	// in:body
	KeyValueItem KeyValueItem `json: "keyValueItem"`
}

// EmptyResponse Boş body
// swagger:response
type EmptyResponse struct {
	// Hata Mesaji
	EmptyBody string `json: "-"`
}
