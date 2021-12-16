package order

import (
	"fmt"
	"oms/core"
	"oms/core/models"
	"time"

	"github.com/go-redis/redis"
)

// locker implementation with Redis
type Locker struct {
	Cache *redis.Client
}

func (s *Locker) Lock(key string) (success bool, err error) {
	res, err := s.Cache.SetNX(key, time.Now().String(), time.Second*15).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

func (s *Locker) Unlock(key string) error {
	return s.Cache.Del(key).Err()
}

func SyncOmsOrderFromUms(warehouseId string) {

	fmt.Printf("Start sync %s!\n", warehouseId)
	var stockWarehouse models.StockWarehouse
	core.App.GetDbByKey("mia_mirror").First(&stockWarehouse, warehouseId)
	fmt.Println(stockWarehouse)
	// t := time.NewTicker(time.Millisecond * 100)
	// c := make(chan struct{})
	// time.AfterFunc(time.Second*5, func() {
	// 	close(c)
	// })

	// for {
	// 	select {
	// 	case <-t.C:
	// 		db.First(&)
	// 		fmt.Print(".")
	// 	case <-c:
	// 		fmt.Println()
	// 		return
	// 	}
	// }
}
