package crontab

import (
	"github.com/robfig/cron"
	"morris/im/helper"
	"morris/im/models"
)

// InitCronTab 启动定时任务
func InitCronTab() {
	crontab := cron.New()

	//秒 分钟 小时 天数  月 星期几
	_ = crontab.AddFunc("0 */1 * * * *", func() {
		helper.Logger.Info("定时任务触发")
	})

	_ = crontab.AddFunc("0 */1 * * * *", func() {
		userModel := new(models.UserModel)
		err := userModel.FindUserByMobile("110")
		if err != nil {
			helper.Logger.Info("定时任务查询数据库失败" + err.Error())
			return
		}

		helper.Logger.Info("110 还没有入住本系统")
	})

	crontab.Start()
}
