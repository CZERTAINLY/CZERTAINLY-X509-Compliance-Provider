package attributes

type AttributeType string

const (
	STRING  AttributeType = "string"
	INTEGER AttributeType = "integer"
	BOOLEAN AttributeType = "boolean"
)

type (
	Attributes struct {
		UUID        string        `json:"uuid,omitempty"`
		Name        string        `json:"name,omitempty"`
		Label       string        `json:"label,omitempty"`
		Type        AttributeType `json:"type,omitempty"`
		Required    bool          `json:"required,omitempty"`
		ReadOnly    bool          `json:"readOnly,omitempty"`
		Editable    bool          `json:"editable,omitempty"`
		Visible     bool          `json:"visible,omitempty"`
		MultiSelect bool          `json:"multiSelect,omitempty"`
		Description bool          `json:"description,omitempty"`
		Content     interface{}   `json:"content,omitempty"`
		List        bool          `json:"list"`
	}

	BaseAttributeContent struct {
		Value interface{} `json:"value"`
	}
)
