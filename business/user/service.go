package user

type service struct {
	repository Repository
}

func (s service) Register(email string, password string) (string, error) {
	return s.repository.Register(email, password)
}

func (s service) Login(email string, password string) (string, error) {
	panic("implement me")
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) FindUserByEmailAndPassword(email string, password string) (*User, error) {
	return s.repository.FindUserByEmailAndPassword(email, password)
}
