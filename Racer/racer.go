package racer

import (
	"fmt"
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

//Ping a url
func Ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

const tenSecondTimeout = 10 * time.Second

//Racer between 2 URL request
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

//ConfigurableRacer between 2 URL request with custom timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-Ping(a):
		return a, nil
	case <-Ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %q and %q", a, b)
	}
}
