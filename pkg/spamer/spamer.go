package spamer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	zlog "github.com/rs/zerolog/log"

	"github.com/zvlb/log-spammer/pkg/config"
)

func Spamer(cfg config.Config) {
	sleepMS := getMSSleepCount(cfg.PerSecond)

	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	counter := 1

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			return
		default:
			zlog.Info().Msg(createMessage(cfg.SizeFrom, cfg.SizeTo))
			time.Sleep(time.Duration(sleepMS * 1000000))

			if cfg.Limit != -1 && counter >= cfg.Limit {
				cansel()
			}

			counter++
		}
	}

}

func getMSSleepCount(perSecond int) int {
	return 1000 / perSecond
}

func createMessage(lengthMin, lengthMax int) string {
	if lengthMin == lengthMax {
		return createText(lengthMax)
	}
	length := randInt(lengthMin, lengthMax)

	return createText(length)

}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func createText(size int) string {
	var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)

}
