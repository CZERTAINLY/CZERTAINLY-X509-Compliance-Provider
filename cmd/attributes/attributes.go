package attributes

type (
	Attributes struct {
		UUID        string      `json:"uuid,omitempty"`
		Name        string      `json:"name,omitempty"`
		Label       string      `json:"label,omitempty"`
		Type        string      `json:"type,omitempty"`
		Required    bool        `json:"required,omitempty"`
		ReadOnly    bool        `json:"readOnly,omitempty"`
		Editable    bool        `json:"editable,omitempty"`
		Visible     bool        `json:"visible,omitempty"`
		MultiValue  bool        `json:"multiValue,omitempty"`
		Description bool        `json:"description,omitempty"`
		Value       interface{} `json:"value,omitempty"`
	}
)
