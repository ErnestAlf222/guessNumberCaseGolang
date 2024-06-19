package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Es de mañana")

	} else if hora < 17 {
		fmt.Println("Es tarde")
	} else {
		fmt.Println("Es noche")

	}

}
