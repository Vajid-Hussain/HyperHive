package cron_auth_server

import (
	repository_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository"
	"github.com/robfig/cron"
)

type CronJob struct {
	userRepository repository_auth_server.UserRepository
}

func NewCronJob(userRepo repository_auth_server.UserRepository) *CronJob {
	return &CronJob{userRepository: userRepo}
}

func (c *CronJob) DeleteExpiredStatus() {
	newCron := cron.New()

	newCron.AddFunc(   , c.DeleteExpiredStatus())
}
