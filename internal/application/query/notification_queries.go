package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
)

type FetchNotificationByID struct {
	ID string
}

type FetchNotificationByIDHandler struct {
	repo notification.NotificationRepositoryInterface
}

func NewFetchNotificationByIDHandler(repo notification.NotificationRepositoryInterface) *FetchNotificationByIDHandler {
	return &FetchNotificationByIDHandler{repo: repo}
}

func (h *FetchNotificationByIDHandler) Handle(ctx context.Context, q FetchNotificationByID) (*notification.Notification, error) {
	return h.repo.GetByID(q.ID)
}

type FetchNotificationList struct{}

type FetchNotificationListHandler struct {
	repo notification.NotificationRepositoryInterface
}

func NewFetchNotificationListHandler(repo notification.NotificationRepositoryInterface) *FetchNotificationListHandler {
	return &FetchNotificationListHandler{repo: repo}
}

func (h *FetchNotificationListHandler) Handle(ctx context.Context, q FetchNotificationList) ([]*notification.Notification, error) {
	return h.repo.List()
}
