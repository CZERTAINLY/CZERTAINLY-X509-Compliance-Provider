package rules

type Service interface {
	// GetRules returns the list of rules available in the connector.
	GetRules(kind string, certificateType []string) ([]Response, error)

	// GetGroups returns the list of groups available in the connector.
	GetGroups(kind string) ([]GroupResponse, error)

	// GetGroupRules returns the list of rules available in the group.
	GetGroupDetails(uuid string, kind string) ([]Response, error)
}
