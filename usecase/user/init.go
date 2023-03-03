package user

import userRepository "lelo-user/repository/user"

type UserUsecaseModule struct {
	Repo userRepository.UserRepository
}

func NewUserusecase(repo userRepository.UserRepository) *UserUsecaseModule {
	return &UserUsecaseModule{
		Repo: repo,
	}
}
