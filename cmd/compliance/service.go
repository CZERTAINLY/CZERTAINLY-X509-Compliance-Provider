package compliance

type Service interface {
	// ComplianceCheck returns the compliance status for the given request.
	//Accepts the certificate and the kind of the certificate and rules to be applied for them
	//Returns the compliance status of the certificate
	ComplianceCheck(kind string, request Request) (Response, error)
}
