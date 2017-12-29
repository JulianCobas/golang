package main

import (
	"fmt"

	"github.com/JulianCobas/golang/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Comedme el nabo"))
	fmt.Println("Fdo: " + stringutil.MyName)
}
