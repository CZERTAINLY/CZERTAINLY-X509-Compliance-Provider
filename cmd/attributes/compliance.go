package attributes

type (
	RequestAttributes struct {
		Name  string      `json:"name"`
		UUID  string      `json:"uuid"`
		Value interface{} `json:"value"`
	}
)
