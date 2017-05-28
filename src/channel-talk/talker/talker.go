package talker

import (
	"fmt"
	"math"
	"time"
)

/**
 * EchoBack - Repeat back like an echo
 * @type {[type]}
 */
func EchoBack(c chan string) {
	for {
		toRepeat := <-c
		fmt.Println("ECHO: ", toRepeat)
		time.Sleep(time.Second * 1)

		for i := 0; i < 1000000000; i++ {
			omg := math.Sqrt(float64(i))
			c <- fmt.Sprintf("%v", omg)
		}
	}
}
