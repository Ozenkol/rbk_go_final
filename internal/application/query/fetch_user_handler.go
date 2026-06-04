package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type UserFilter struct {
	Email string
	CreatedAfter string
	CreatedBefore string

}

type FetchUserQuery struct {
	UserID string
	Filter UserFilter
	Pagination application_shared.Pagination
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
