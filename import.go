package main

import (
	"learn-go/helper"
)

func main() {
	result := helper.SayHello("Ibnu")
	println(result)

	println(helper.Application)
	println(helper.SayGoodbye("Ibnu"))
}
