package initialize

import (
	"fmt"
	"github.com/my-gin-web/utils/dbInstance"

	"github.com/my-gin-web/config"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/utils"
)

func Timer() {
	if global.Config.Timer.Start {
		for i := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				_, _ = global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {
					err := utils.ClearTable(dbInstance.SelectConn(), detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})

			}(global.Config.Timer.Detail[i])
		}
	}
}
