package health

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) GetHealth() Health {
	return Health{
		Status:      "ok",
		Description: "OK",
	}
}
