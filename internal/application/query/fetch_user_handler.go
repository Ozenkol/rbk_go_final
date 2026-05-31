package query

import "github.com/Ozenkol/rbk-go-final/internal/domain/user"

type FetchUserQuery struct {
	UserID string
}

type FetchUserHandler struct {
	userRepo user.UserRepositoryInterface
}

func NewFetchUserHandler(userRepo user.UserRepositoryInterface) *FetchUserHandler {
	return &FetchUserHandler{
		userRepo: userRepo,
	}
}

func (h *FetchUserHandler) Handle(query FetchUserQuery) (*user.User, error) {
	return h.userRepo.GetByID(query.UserID)
}