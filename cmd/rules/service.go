package rules

type Service interface {
	GetRules(kind string, certificateType []string) ([]Response, error)
	GetGroups(kind string) ([]GroupResponse, error)
	GetGroupDetails(uuid string, kind string) ([]Response, error)
}
