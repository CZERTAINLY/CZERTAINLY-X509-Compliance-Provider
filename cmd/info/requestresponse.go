package info

type (
	EndPointsInfo struct {
		Name    string `json:"name"`
		Context string `json:"context"`
		Method  string `json:"method"`
	}

	Response struct {
		FunctionGroupCode string          `json:"functionGroupCode"`
		Kinds             []string        `json:"kinds"`
		EndPoints         []EndPointsInfo `json:"endPoints"`
	}
)
