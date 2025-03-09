package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	t := time.Now()
	fmt.Printf("[INFO] %s: %s\n", t.Format("2023-12:12 23:23"), msg)

}
