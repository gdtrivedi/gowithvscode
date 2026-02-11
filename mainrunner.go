package main

import (
	"github.com/gdtrivedi/gowithvscode/dequegammazero"
	"github.com/gdtrivedi/gowithvscode/dequepkg"
	"github.com/gdtrivedi/gowithvscode/helloworld"
)

func main() {
	pkg_helloworld()
	// pkg_dequqpkg()
	// pkg_dequegammazero()
}

func pkg_dequegammazero() {
	dequegammazero.DequeGammazeroTest()
}

func pkg_dequqpkg() {
	dequepkg.DequeTest()
}

func pkg_helloworld() {
	helloworld.HelloWorld()
}
