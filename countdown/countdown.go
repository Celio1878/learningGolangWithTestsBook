package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	const finalWord = "GO!"
	const countdownStart = 3

	for i := countdownStart; i > 0; i-- {
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			fmt.Println(err)
		}
		sleeper.Sleep()
	}

	_, err := fmt.Fprint(out, finalWord)
	if err != nil {
		fmt.Println(err)
	}
}
