package main

import "fmt"

func main() {
	b := Bird{}
	b.colour = "black"
	b.weight = 228
	b.tail.length = 322
	b.tail.width = 233
	b.wing.length = 1337
	b.wing.square = 7331
	fmt.Println(b)

}

type (
	Wing struct {
		square float32
		length float32
	}

	Tail struct {
		width  float32
		length float32
	}
)

type (
	Bird struct {
		tail   Tail
		wing   Wing
		colour string
		weight float32
	}
)
