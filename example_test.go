package timereplace_test

import (
	"fmt"
	"time"

	"github.com/koron-go/timereplace"
)

func ExampleReplace() {
	origtime := time.Date(2006, 1, 2, 3, 4, 5, 999999999, time.UTC)
	fmt.Println(origtime.Format(time.RFC3339Nano))

	modtime := timereplace.Replace(origtime, timereplace.Values{
		Year:  new(2026),
		Month: new(time.Month(2)),
		Day:   new(19),
	})
	fmt.Println(modtime.Format(time.RFC3339Nano))

	// Output:
	// 2006-01-02T03:04:05.999999999Z
	// 2026-02-19T03:04:05.999999999Z
}
