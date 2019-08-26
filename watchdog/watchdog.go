package watchdog

import (
	"os"
	"time"

	"github.com/anderskvist/GoHelpers/log"
)

var timer time.Time

// Poke the watchdog
func Poke() {
	log.Debug("WatchDog: GRRR!")
	timer = time.Now()
}

// Activate the watchdog
func Activate(interval int) {
	Poke()
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for ; true; <-ticker.C {
		age := time.Since(timer)
		log.Debugf("WatchDog: Woof! (%.2fs)", age.Seconds())
		if age/1000000000 > 15 {
			log.Debugf("WatchDog: zzzzZzZZZzzzzZz... (%.2fs)", age.Seconds())
			os.Exit(1)
		}
	}
}
