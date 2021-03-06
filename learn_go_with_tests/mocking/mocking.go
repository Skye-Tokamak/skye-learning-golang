package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}
type CountDownSpyer struct {
	Calls []string
}

func (s *CountDownSpyer) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountDownSpyer) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}
func Countdown(out io.Writer, sleeper Sleeper) {
	count := 3
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "GO")
}
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
