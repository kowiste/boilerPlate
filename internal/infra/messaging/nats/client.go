package nats

import (
	"context"
	"fmt"
	"time"

	"ddd/pkg/config"
	"ddd/shared/logger"

	"github.com/nats-io/nats.go"
)

type Client struct {
	conn        *nats.Conn
	js          nats.JetStreamContext
	logger      logger.Logger
	// natsWrapper *trace.NatsWrapper
	// tracer      *trace.Tracer
}

func NewClient(ctx context.Context, cfg *config.Config, logger logger.Logger) (*Client, error) {
	opts := []nats.Option{
		nats.Name(cfg.App.Name),
		nats.Timeout(cfg.NATS.Timeout),
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.Error(ctx, err, "NATS disconnected: %v", map[string]interface{}{})
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Info(ctx, "NATS reconnected", map[string]interface{}{})
		}),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			logger.Error(ctx, err, "NATS error: %v", map[string]interface{}{})
		}),
	}

	nc, err := nats.Connect(cfg.NATS.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	return &Client{
		conn:        nc,
		js:          js,
		logger:      logger,
		// natsWrapper: trace.NewNatsWrapper(nc, cfg.App.Name),
		// tracer:      trace.NewTracer("nats-client"),
	}, nil
}

// func (c *Client) Close() {
// 	if c.conn != nil {
// 		c.conn.Close()
// 	}
// }

// func (c *Client) Publish(ctx context.Context, subject string, data []byte) error {
// 	return c.natsWrapper.Publish(ctx, subject, data)
// }

// func (c *Client) Subscribe(subject string, handler func(ctx context.Context, msg *nats.Msg)) (*nats.Subscription, error) {
// 	return c.natsWrapper.Subscribe(subject, handler)
// }

// // QueueSubscribe for load balancing among subscribers
// func (c *Client) QueueSubscribe(subject, queue string, handler func(ctx context.Context, msg *nats.Msg)) (*nats.Subscription, error) {
// 	return c.conn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
// 		ctx, span := c.tracer.StartSpan(context.Background(), "nats.queue.subscribe")
// 		defer span.End()

// 		span.SetAttributes(
// 			attribute.String("messaging.system", "nats"),
// 			attribute.String("messaging.destination", subject),
// 			attribute.String("messaging.nats.queue", queue),
// 		)

// 		handler(ctx, msg)
// 	})
// }

// // StreamSubscribe for persistent messaging
// func (c *Client) StreamSubscribe(subject string, handler func(ctx context.Context, msg *nats.Msg)) (*nats.Subscription, error) {
// 	stream := &nats.StreamConfig{
// 		Name:     "MEASUREMENTS",
// 		Subjects: []string{subject},
// 		Storage:  nats.FileStorage,
// 	}

// 	_, err := c.js.AddStream(stream)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create stream: %w", err)
// 	}

// 	return c.js.Subscribe(subject, func(msg *nats.Msg) {
// 		ctx, span := c.tracer.StartSpan(context.Background(), "nats.stream.subscribe")
// 		defer span.End()

// 		span.SetAttributes(
// 			attribute.String("messaging.system", "nats"),
// 			attribute.String("messaging.destination", subject),
// 			attribute.String("messaging.nats.stream", stream.Name),
// 		)

// 		handler(ctx, msg)
// 	})
// }
