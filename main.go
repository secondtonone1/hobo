package main

import (
	_ "hobo/simple_service"

	_ "hobo/components"

	node "github.com/duanhf2012/origin/node"
)

func main() {
	node.Start()
}
