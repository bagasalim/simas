package socket_system

type Service interface {
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
