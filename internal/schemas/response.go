package schema

type ResponsePayload struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
