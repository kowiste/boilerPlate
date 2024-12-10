package publisher

import (
	"context"

	"ddd/internal/infra/messaging/nats"
	"ddd/shared/logger"
)

type Publisher struct {
	natsClient *nats.Client
	logger     *logger.Logger
}

func NewPublisher(natsClient *nats.Client, logger *logger.Logger) *Publisher {
	return &Publisher{
		natsClient: natsClient,
		logger:     logger,
	}
}

func (p *Publisher) PublishAssetCreated(ctx context.Context, orgID string, asset interface{}) error {
	subject := nats.SubjectForOrg(nats.SubjectAssetCreated, orgID)
	return p.publish(ctx, subject, asset)
}

func (p *Publisher) PublishMeasure(ctx context.Context, orgID string, measure interface{}) error {
	subject := nats.SubjectForOrg(nats.SubjectMeasureCreated, orgID)
	return p.publish(ctx, subject, measure)
}

func (p *Publisher) PublishAlert(ctx context.Context, orgID string, alert interface{}) error {
	subject := nats.SubjectForOrg(nats.SubjectNotificationAlert, orgID)
	return p.publish(ctx, subject, alert)
}

func (p *Publisher) PublishNotification(ctx context.Context, orgID string, notification interface{}) error {
	subject := nats.SubjectForOrg(nats.SubjectNotificationInfo, orgID)
	return p.publish(ctx, subject, notification)
}

func (p *Publisher) publish(ctx context.Context, subject string, data interface{}) error {
	// payload, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }

	// return p.natsClient.Publish(ctx, subject, payload)
	return nil
}
