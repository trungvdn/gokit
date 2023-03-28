package chaptereight

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type PendingUserNotifications map[int][]*Notification
type Notification struct {
	Content string
	UserId  int
}

func sendUserBatchNotificationsEmail(userId int, notifications []*Notification) {
	fmt.Printf("Sending email to user with userId %d for pending notifications %v\n", userId, notifications)
}

func handlePendingUsersNotifications(pendingNotifications PendingUserNotifications, handler func(userId int, notifications []*Notification)) {
	for userId, notifications := range pendingNotifications {
		handler(userId, notifications)
		delete(pendingNotifications, userId)
	}
}

func collectNewUsersNotifications(notifications PendingUserNotifications) {
	randomNotifications := getRandomNotifications()
	if len(randomNotifications) > 0 {
		notifications[randomNotifications[0].UserId] = randomNotifications
	}
}

func getRandomNotifications() (notifications []*Notification) {
	rand.Seed(time.Now().UnixNano())
	userId := rand.Intn(100-10+1) + 10
	numOfNotifications := rand.Intn(5-0+1) + 0
	fmt.Printf("numOfNotifications %v\n", numOfNotifications)
	for i := 0; i < numOfNotifications; i++ {
		notifications = append(notifications, &Notification{Content: gofakeit.Paragraph(1, 2, 10, " "), UserId: userId})
	}

	return
}
