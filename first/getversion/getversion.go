package getversion

import (
	"fmt"
	"runtime"
)

//Goversion return the version of go
func Goversion() {
	fmt.Println(runtime.Version())

}
