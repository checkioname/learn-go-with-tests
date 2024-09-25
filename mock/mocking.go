package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


type Sleeper interface {
  Sleep()
}

// Mock sleeper (will be used for tests)
type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}

// real sleeper
type DefaultSleeper struct {} 

func (s *DefaultSleeper) Sleep() {
  time.Sleep(time.Second*1)
}


func Countdown(out io.Writer, start int, sleeper Sleeper) {
  for i := start; i >0; i -- {
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }
}



// Extending sleeper (will be configurable)

type ConfigurableSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
  c.sleep(c.duration)
}

type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}


func main(){
  // passing real sleeper
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout, 3, sleeper)
}



