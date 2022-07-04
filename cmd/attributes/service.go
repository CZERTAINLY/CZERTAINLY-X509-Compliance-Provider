package attributes

type Service interface {
	// GetAttributes returns the list of attributes for a kind. Since the compliance provider connector
	// does not have any attributes, this function returns an empty list.
	GetAttributes() ([]interface{}, error)
	ValidateAttributes() ([]interface{}, error)
}
