package progressbar

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Render render a progress bar and return set percent func and stop func.
func Render() (func(p float64), func()) {
	var (
		done bool
		pg   float64
	)
	go func() {
		for i, c := 0, 60; pg <= 1; i++ {
			if done {
				return
			}
			if pg == 1 {
				done = true
			}
			// Render text and padding.
			str := fmt.Sprintf("\r%.2f%%%s", pg*1e+2, strings.Repeat(" ", 10))
			// Render bar.
			n := int(pg * 1e+2 / (float64(100) / float64(c)))
			str += "[" + strings.Repeat("█", n)
			if n < c {
				str += strings.Repeat("-", c-n)
			}
			str += "]"
			// Render spinner.
			str += fmt.Sprintf(" %c", `-\|/`[i%4])
			io.WriteString(os.Stdout, str)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return func(p float64) {
			pg = p
		}, func() {
			done = true
		}
}
