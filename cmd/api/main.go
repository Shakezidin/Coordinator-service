package main

import (
	admindi "github.com/Shakezidin/pkg/admin/di"
	coordinatordi "github.com/Shakezidin/pkg/coordinator/di"
)


func main() {
	admindi.Init()
	coordinatordi.Init()
}