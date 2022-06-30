package attributes

type (
	RequestAttributes struct {
		Name    string      `json:"name"`
		UUID    string      `json:"uuid"`
		Content interface{} `json:"content"`
	}
)
