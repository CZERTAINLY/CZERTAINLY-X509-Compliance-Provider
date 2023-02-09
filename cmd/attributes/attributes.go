package attributes

type AttributeType string
type AttributeContentType string

// AttributeType constants. These are the types of attributes that can be used in the rules.
const (
	STRING  AttributeContentType = "string"
	INTEGER AttributeContentType = "integer"
	BOOLEAN AttributeContentType = "boolean"
)

const (
	DATA  AttributeType = "data"
	INFO  AttributeType = "group"
	GROUP AttributeType = "info"
)

// Attribute is a single attribute that can be used in the rules.
type (
	Attributes struct {
		UUID        string                 `json:"uuid,omitempty"`
		Name        string                 `json:"name,omitempty"`
		Type        AttributeType          `json:"type,omitempty"`
		ContentType AttributeContentType   `json:"content_type,omitempty"`
		Description string                 `json:"description,omitempty"`
		Properties  AttributeProperties    `json:"properties,omitempty"`
		Content     []BaseAttributeContent `json:"content,omitempty"`
	}

	AttributeProperties struct {
		Label       string `json:"label,omitempty"`
		Required    bool   `json:"required,omitempty"`
		ReadOnly    bool   `json:"readOnly,omitempty"`
		Editable    bool   `json:"editable,omitempty"`
		Visible     bool   `json:"visible,omitempty"`
		MultiSelect bool   `json:"multi,omitempty"`
		List        bool   `json:"list"`
	}

	// BaseAttributeContent for the values of attributes
	BaseAttributeContent struct {
		Data interface{} `json:"data,omitempty"`
	}

	RequestAttributes struct {
		Name    string                 `json:"name"`
		UUID    string                 `json:"uuid"`
		Content []BaseAttributeContent `json:"content"`
	}
)
