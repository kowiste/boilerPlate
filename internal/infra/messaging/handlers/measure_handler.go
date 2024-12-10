package handlers

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/nats-io/nats.go"
// 	measure "ddd/internal/features/measure/domain"
// 	"ddd/shared/logging"
// )

// type MeasureHandler struct {
// 	measureService measure.Service
// 	logger         *logging.Logger
// }

// func NewMeasureHandler(measureService measure.Service, logger *logging.Logger) *MeasureHandler {
// 	return &MeasureHandler{
// 		measureService: measureService,
// 		logger:        logger,
// 	}
// }

// func (h *MeasureHandler) HandleNewMeasure(ctx context.Context, msg *nats.Msg) {
// 	var measure measure.Measure
// 	if err := json.Unmarshal(msg.Data, &measure); err != nil {
// 		h.logger.Errorf("Failed to unmarshal measure: %v", err)
// 		return
// 	}

// 	if err := h.measureService.ProcessMeasure(ctx, measure); err != nil {
// 		h.logger.Errorf("Failed to process measure: %v", err)
// 		return
// 	}

// 	msg.Ack()
// }

// internal/infrastructure/messaging/handlers/notification_handler.go
