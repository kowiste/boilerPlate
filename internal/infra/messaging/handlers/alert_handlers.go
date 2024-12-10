package handlers

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/nats-io/nats.go"
// 	"ddd/shared/logging"
// )

// type AlertHandler struct {
// 	alertService alert.Service
// 	logger       *logging.Logger
// }

// func NewAlertHandler(alertService alert.Service, logger *logging.Logger) *AlertHandler {
// 	return &AlertHandler{
// 		alertService: alertService,
// 		logger:       logger,
// 	}
// }

// func (h *AlertHandler) HandleAlert(ctx context.Context, msg *nats.Msg) {
// 	var alert alert.Alert
// 	if err := json.Unmarshal(msg.Data, &alert); err != nil {
// 		h.logger.Errorf("Failed to unmarshal alert: %v", err)
// 		return
// 	}

// 	if err := h.alertService.ProcessAlert(ctx, alert); err != nil {
// 		h.logger.Errorf("Failed to process alert: %v", err)
// 		return
// 	}

// 	msg.Ack()
// }
