package nats

import "fmt"

const (
	// Asset related subjects
	SubjectAssetCreated = "asset.created"
	SubjectAssetUpdated = "asset.updated"
	SubjectAssetDeleted = "asset.deleted"

	// Measure related subjects
	SubjectMeasureCreated    = "measure.created"
	SubjectMeasureProcessed  = "measure.processed"
	SubjectMeasureAggregated = "measure.aggregated"

	// Dashboard related subjects
	SubjectDashboardCreated = "dashboard.created"
	SubjectDashboardUpdated = "dashboard.updated"
	SubjectWidgetUpdated    = "widget.updated"

	// Notification subjects
	SubjectNotificationAlert = "notification.alert"
	SubjectNotificationInfo  = "notification.info"

	// Stream names
	StreamMeasurements = "MEASUREMENTS"
	StreamAssets       = "ASSETS"
	StreamAlerts       = "ALERTS"
)

// SubjectForOrg returns organization-specific subject
func SubjectForOrg(subject, orgID string) string {
	return fmt.Sprintf("%s.%s", subject, orgID)
}