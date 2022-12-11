package http

// Response stores data of user request result.
type Response struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
