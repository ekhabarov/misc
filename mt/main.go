package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())

	right := 0
	wrong := 0

	start := time.Now()

	defer func() {
		s := time.Since(start).Seconds()

		fmt.Printf("Equations done: %d, correct: %d, incorrect: %d\n",
			right+wrong, right, wrong)
		fmt.Printf("Withing %.2f seconds\n", s)
		fmt.Printf("Your speed is %.2f equations per second\n", s/float64(right))
	}()

	for {
		a := rand.Int31n(12)
		b := rand.Int31n(12)
		c := a * b

		fmt.Printf("%dx%d= ", a, b)

		r, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("err:", err)
		}

		in := r[:len(r)-1]

		if string(in) == "q" {
			return
		}

		ir, err := strconv.Atoi(string(in))
		if err != nil {
			fmt.Println("it wasn't a number")
		}

		if ir == int(c) {
			fmt.Println("+")
			right++
		} else {
			fmt.Println("-")
			wrong++
		}
	}
}
