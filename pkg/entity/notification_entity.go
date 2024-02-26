package entity

type NotificationPayload struct {
	NotificationType string
	Subject          string
	Body             string
	Target           string
}
