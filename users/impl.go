package users

type UserServiceImpl struct {
	repository *UserRepositoryImpl
}

func NewUserService(repository *UserRepositoryImpl) UserService {
	return &UserServiceImpl{repository: repository}
}

func (s *UserRepositoryImpl) CreateUserService() {

}
