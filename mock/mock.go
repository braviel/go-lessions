package mock

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//Sleeper bring sleep time out of Countdown
type Sleeper interface {
	Sleep()
}

//ConfigurableSleeper sleep with specific duration
type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepImpl func(d time.Duration)
}

//Sleep interface impl
func (c *ConfigurableSleeper) Sleep() {
	c.SleepImpl(c.Duration)
}

//Countdown number
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }
	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }
	// sleeper.Sleep()
	// fmt.Fprint(out, finalWord)
}
