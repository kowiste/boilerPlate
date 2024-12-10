package handlers

// import (
// 	"context"
// 	"ddd/shared/logging"

// 	"github.com/nats-io/nats.go"
// )

// type NotificationHandler struct {
// 	websocketHub *websocket.Hub
// 	logger      *logging.Logger
// }

// func NewNotificationHandler(hub *websocket.Hub, logger *logging.Logger) *NotificationHandler {
// 	return &NotificationHandler{
// 		websocketHub: hub,
// 		logger:      logger,
// 	}
// }

// func (h *NotificationHandler) HandleNotification(ctx context.Context, msg *nats.Msg) {
// 	// Broadcast to all connected clients
// 	h.websocketHub.Broadcast(msg.Data)
// 	msg.Ack()
// }

