package packets

type Error_response struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
