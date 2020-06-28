package paninc_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanincVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("enter defer")
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recovered from ", err)
	// 	}
	// }()
	fmt.Println("Start")
	panic(errors.New("Something wrong."))
	os.Exit((-1)) //not enter defer
}
