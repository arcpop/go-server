package trace

import (
	"fmt"
)


func trace_enter(f string) {
	fmt.Println("Entering ", f);
}
func trace_exit(f string) {
	fmt.Println("Exiting ", f);
}
