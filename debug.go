package logprint

import (
	"fmt"
	"time"
)

func Debug(msg interface{}) {
	t := time.Now()
	fmt.Printf("[DEBUG] %s: %s\n", t.Format("2023-12:12 23:23"), msg)

}
