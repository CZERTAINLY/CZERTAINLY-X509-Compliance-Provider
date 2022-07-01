package attributes

type AttributeType string

// AttributeType constants. These are the types of attributes that can be used in the rules.
const (
	STRING  AttributeType = "string"
	INTEGER AttributeType = "integer"
	BOOLEAN AttributeType = "boolean"
)

// Attribute is a single attribute that can be used in the rules.
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

	// BaseAttributeContent for the values of attributes
	BaseAttributeContent struct {
		Value interface{} `json:"value"`
	}

	RequestAttributes struct {
		Name    string      `json:"name"`
		UUID    string      `json:"uuid"`
		Content interface{} `json:"content"`
	}
)
