package main

import "internal/singleton"

func main() {
	u := singleton.Instance()
	u.Setname("who am i")
	u.Say()

	v := singleton.Instance()
	v.Say()
}
