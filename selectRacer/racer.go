package selectRacer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			return
		}

		close(channel)
	}()

	return channel
}
