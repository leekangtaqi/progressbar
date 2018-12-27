package progressbar

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// NewRender receive duration for throttle, returns render function to rerender
// progress bar.
func NewRender(rate time.Duration) func(percent float64) {
	var (
		done bool
		i    int
		throttle = time.Tick(rate)
	)

	return func(pc float64) {
		<-throttle
		if pc > 1 || pc < 0 {
			panic(fmt.Errorf("percent %f invalid", pc))
		}
		if done {
			return
		}
		if pc == 1 {
			done = true
		}
		// Render text and padding.
		str := fmt.Sprintf("\r%.2f%%%s", pc*1e+2, strings.Repeat(" ", 10))
		// Render bar.
		n := int(pc * 1e+2 / (float64(100) / float64(60)))
		str += "[" + strings.Repeat("█", n)
		if n < 60 {
			str += strings.Repeat("-", 60-n)
		}
		// Render spinner.
		io.WriteString(os.Stdout, str+fmt.Sprintf("] %c", `-\|/`[i%4]))
		i++
	}
}
