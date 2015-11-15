// order project main.go
package main

import (
	"github.com/gernest/utron"
	_ "order/controllers"
	_ "order/models"
)

func main() {
	utron.Run()
}
