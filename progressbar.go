package progressbar

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// NewRender receive custom writer and duration for throttle, returns render
// function to rerender progress bar.
func NewRender(writer io.Writer, rate time.Duration) func(percent float64) error {
	var (
		i        int
		throttle = time.Tick(rate)
	)

	return func(pc float64) error {
		<-throttle
		if pc > 1 || pc < 0 {
			return fmt.Errorf("percent %f invalid", pc)
		}
		// Render text and padding.
		pl := 6 - len(strconv.Itoa(int(pc*1e+2))
		str := fmt.Sprintf("\r%.2f%%%s", pc*1e+2, strings.Repeat(" ", pl))
		// Render bar.
		n := int(pc * 1e+2 / (float64(100) / float64(60)))
		str += fmt.Sprintf("[%s%s", strings.Repeat("█", n), strings.Repeat("-", 60-n))
		// Render spinner.
		io.WriteString(writer, str+fmt.Sprintf("] %c", `-\|/`[i%4]))
		i++
		return nil
	}
}
