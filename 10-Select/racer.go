package selectexample

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)

// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool.
// Since we are closing and not sending anything on the chan, why allocate anything?

// Notice how we have to use make when creating a channel; rather than say var ch chan struct{}.
// When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
//For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
