package health

type Service interface {
	// GetHealth returns the health of the connector. Since the compliance provider connector
	// does not have any real checks, it returns a healthy status.
	GetHealth() Health
}
