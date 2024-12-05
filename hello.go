package main

import "fmt"

// import "go.uber.org/zap"
// import "github.ibm.com/IAM/basiclog"
import "github.ibm.com/sos-ws-scan/greetings"

func main() {
	// fmt.Println(quote.Go())
	// Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
