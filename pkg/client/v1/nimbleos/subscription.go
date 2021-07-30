// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Subscription - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification client.
// Export SubscriptionFields for advance operations like search filter etc.
var SubscriptionFields *SubscriptionStringFields

func init() {
	IDfield := "id"
	SubscriberIdfield := "subscriber_id"
	NotificationTypefield := "notification_type"
	ObjectTypefield := "object_type"
	ObjectIdfield := "object_id"
	Operationfield := "operation"
	EventTargetTypefield := "event_target_type"
	EventSeverityfield := "event_severity"

	SubscriptionFields = &SubscriptionStringFields{
		ID:               &IDfield,
		SubscriberId:     &SubscriberIdfield,
		NotificationType: &NotificationTypefield,
		ObjectType:       &ObjectTypefield,
		ObjectId:         &ObjectIdfield,
		Operation:        &Operationfield,
		EventTargetType:  &EventTargetTypefield,
		EventSeverity:    &EventSeverityfield,
	}
}

type Subscription struct {
	// ID - Identifier for subscription.
	ID *string `json:"id,omitempty"`
	// SubscriberId - Identifier for subscriber (notification client) that this subscription belongs to.
	SubscriberId *string `json:"subscriber_id,omitempty"`
	// NotificationType - This indicates the type of notification being subscribed for.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// ObjectType - The object type that the notification subscriber is interested in. This is relevant for and required only for audit log based notifications.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - The object that the notification subscriber is interested in. Applies only to audit log based notifications.
	ObjectId *string `json:"object_id,omitempty"`
	// Operation - The operation that the notification subscriber is interested in. Applies only to audit log based notifications.
	Operation *NsAuditOperationTypeOrAll `json:"operation,omitempty"`
	// EventTargetType - The kind of events or alerts that the notification subscriber is interested in. Applies only to events based notifications.
	EventTargetType *NsEventTargetTypeOrAll `json:"event_target_type,omitempty"`
	// EventSeverity - The severity of events that the notification subscriber is interested in. Applies only to events based notifications.
	EventSeverity *NsSeverityLevel `json:"event_severity,omitempty"`
}

// Struct for SubscriptionFields
type SubscriptionStringFields struct {
	ID               *string
	SubscriberId     *string
	NotificationType *string
	ObjectType       *string
	ObjectId         *string
	Operation        *string
	EventTargetType  *string
	EventSeverity    *string
}
