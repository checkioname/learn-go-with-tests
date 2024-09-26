package main

import (
	"bytes"
	"testing"
	"time"
)

//Our software needs to print to stdout and we saw how we could
//use Dependency Injection (DI) to facilitate testing this in the DI section.

// Usamos o time.Sleep -> seria interessante mockar isso para otimizar nossos testes


func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  // passing fake sleeper ( using Dependency injection technique)
  sleeper := SpySleeper{}
  Countdown(buffer,3, &sleeper)

  got := buffer.String()
  want := "3\n2\n1\n"

  if got != want {
    t.Errorf("Expected %q but got %q", want, got)
  }


}


func TestConfigurableSleep(t *testing.T) {
  sleepTime := time.Second * 5
  // passing fake sleeper ( using Dependency injection technique)
  spyTime := &SpyTime{}
  sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
  sleeper.Sleep()

  if spyTime.durationSlept != sleepTime {
    t.Errorf("Expected %v but got %v", sleepTime, spyTime.durationSlept)
  }
}
