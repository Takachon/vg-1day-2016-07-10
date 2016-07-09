package request

type (
	Message struct {
		Id       int    `json:id`
		Body     string `json:"body"`
		Username string `json:"user_name"`
	}
)
