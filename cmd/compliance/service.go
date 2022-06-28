package compliance

type Service interface {
	ComplianceCheck(kind string, request Request) (Response, error)
}
