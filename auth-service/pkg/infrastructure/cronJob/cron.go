package cron_auth_server

import (
	"fmt"
	"time"

	interface_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	userRepository interface_repo_auth_server.IUserRepository
	Location       *time.Location
}

func NewCronJob(userRepo interface_repo_auth_server.IUserRepository) *CronJob {
	locationInd, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("error at time place", err)
	}
	return &CronJob{userRepository: userRepo,
		Location: locationInd}
}

func (c *CronJob) StartCronInAuthService() {
	newCron := cron.New()

	newCron.AddFunc("*/15 * * * *", func() {
		c.userRepository.DeleteExpiredStatus(time.Now())
	})

	newCron.AddFunc("0 0 */2 * *", c.userRepository.DeleteUnverifiedUsers)
	newCron.Start()
}
