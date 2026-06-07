package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
)

type CreateNotificationCommand struct {
	Notification *notification.Notification
}

type CreateNotificationHandler struct {
	repo notification.NotificationRepositoryInterface
}

func NewCreateNotificationHandler(repo notification.NotificationRepositoryInterface) *CreateNotificationHandler {
	return &CreateNotificationHandler{repo: repo}
}

func (h *CreateNotificationHandler) Handle(ctx context.Context, cmd CreateNotificationCommand) (*notification.Notification, error) {
	return h.repo.Create(cmd.Notification)
}

type UpdateNotificationCommand struct {
	Notification *notification.Notification
}

type UpdateNotificationHandler struct {
	repo notification.NotificationRepositoryInterface
}

func NewUpdateNotificationHandler(repo notification.NotificationRepositoryInterface) *UpdateNotificationHandler {
	return &UpdateNotificationHandler{repo: repo}
}

func (h *UpdateNotificationHandler) Handle(ctx context.Context, cmd UpdateNotificationCommand) (*notification.Notification, error) {
	return h.repo.Update(cmd.Notification)
}

type DeleteNotificationCommand struct {
	ID string
}

type DeleteNotificationHandler struct {
	repo notification.NotificationRepositoryInterface
}

func NewDeleteNotificationHandler(repo notification.NotificationRepositoryInterface) *DeleteNotificationHandler {
	return &DeleteNotificationHandler{repo: repo}
}

func (h *DeleteNotificationHandler) Handle(ctx context.Context, cmd DeleteNotificationCommand) error {
	return h.repo.Delete(cmd.ID)
}
