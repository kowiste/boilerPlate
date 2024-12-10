package subscriber

import (
	"context"

	// "ddd/internal/infra/messaging/handlers"
	"ddd/internal/infra/messaging/nats"
	"ddd/shared/logger"
)

type Subscriber struct {
	natsClient        *nats.Client
	// measureHandler    *handlers.MeasureHandler
	// notificationHandler *handlers.NotificationHandler
	// alertHandler      *handlers.AlertHandler
	logger           *logger.Logger
}

func NewSubscriber(
	natsClient *nats.Client,
	// measureHandler *handlers.MeasureHandler,
	// notificationHandler *handlers.NotificationHandler,
	// alertHandler *handlers.AlertHandler,
	logger *logger.Logger,
) *Subscriber {
	return &Subscriber{
		natsClient:        natsClient,
		// measureHandler:    measureHandler,
		// notificationHandler: notificationHandler,
		// alertHandler:      alertHandler,
		logger:           logger,
	}
}

func (s *Subscriber) Start(ctx context.Context) error {
	// Subscribe to measure updates with queue group for load balancing
	// if _, err := s.natsClient.QueueSubscribe(
	// 	nats.SubjectMeasureCreated,
	// 	"measure-processors",
	// 	s.measureHandler.HandleNewMeasure,
	// ); err != nil {
	// 	return err
	// }

	// Subscribe to notifications for broadcasting to websocket clients
	// if _, err := s.natsClient.Subscribe(
	// 	nats.SubjectNotificationInfo,
	// 	s.notificationHandler.HandleNotification,
	// ); err != nil {
	// 	return err
	// }

	// // Subscribe to alerts with persistent storage
	// if _, err := s.natsClient.StreamSubscribe(
	// 	nats.SubjectNotificationAlert,
	// 	s.alertHandler.HandleAlert,
	// ); err != nil {
	// 	return err
	// }

	return nil
}