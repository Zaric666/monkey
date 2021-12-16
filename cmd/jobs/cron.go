package jobs

import (
	"oms/core"
	"oms/cron/order"
	"os"

	"github.com/go-redis/redis"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	StartCmd = &cobra.Command{
		Use:     "cron",
		Example: "oms cron",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {

}

func run() {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	l := &order.Locker{Cache: c}

	// Make locker available for the cron jobs
	gocron.SetLocker(l)

	args := os.Args[1:]

	dsn := "houduan_zhangli:Zl123@mia.com@tcp(10.5.96.80:3306)/mia_mirror?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	core.App.SetDb("mia_mirror", db)

	gocron.Every(10).Seconds().Lock().Do(order.SyncOmsOrderFromUms, args[1])
	<-gocron.Start()
}
