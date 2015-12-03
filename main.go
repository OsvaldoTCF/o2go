// order project main.go
package main

import (
	_ "github.com/OsvaldoTCF/order2go/controllers"
	_ "github.com/OsvaldoTCF/order2go/models"
	"github.com/gernest/utron"
)

func main() {
	utron.Run()
}
