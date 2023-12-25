package app

type service struct {
}

func (s *service) NewService() *service {
	return &service{}

}
