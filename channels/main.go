package main

func product(x, y int, ch chan int) {
	p := x * y
	ch <- p
}

func main() {
	ch := make(chan int, 2)

	go product(4, 5, ch)
	go product(6, 8, ch)
	go product(3, 3, ch)
	go product(4, 7, ch)

	x, y := <-ch, <-ch
	println(x, y)

	a, b := <-ch, <-ch
	println(a, b)
}
