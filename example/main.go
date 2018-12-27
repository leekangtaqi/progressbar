package main

import (
	"os"
	"time"

	"github.com/leekangtaqi/progressbar/v1"
)

func main() {
	render := NewRender(os.Stdout, time.Millisecond * 100)
	render(0.012322)
	render(0.022322)
	render(0.035322)
	render(0.123522)
	render(0.282322)
	render(0.302322)
	render(0.312322)
	render(0.352322)
	render(0.412322)
	render(0.422322)
	render(0.462322)
	render(0.472322)
	render(0.512322)
	render(0.59322)
	render(0.662322)
	render(0.772322)
	render(0.802322)
	render(0.902322)
	render(0.982322)
	render(0.992322)
	render(1)
}
