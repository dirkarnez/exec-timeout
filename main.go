// package main

// import (
// 	"context"
// 	"fmt"
// 	"os/exec"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	fmt.Println("Start")
// 	cmd := exec.CommandContext(ctx, "sleep", "5") //os.Args[0], os.Args[1:]...)

// 	_, err := cmd.CombinedOutput()

// 	if ctx.Err() == context.DeadlineExceeded {
// 		fmt.Println("Command was killed")
// 	}

//		if err != nil {
//			fmt.Println("If the command was killed, err will be \"signal: killed\"")
//			fmt.Println("If the command wasn't killed, it contains the actual error, e.g. invalid command")
//		} else {
//			fmt.Println("Done")
//		}
//	}
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <duration_in_seconds> <command> [<args>...]")
		return
	}

	// Parse duration from the first command-line argument
	duration, err := strconv.Atoi(os.Args[1])
	if err != nil || duration <= 0 {
		fmt.Println("Please provide a valid positive integer for duration.")
		return
	}

	// Get the command and its arguments
	command := os.Args[2]
	args := os.Args[3:] // Remaining arguments are command arguments

	fmt.Printf("Starting %s...%v\n", command, args)

	// // Start the command in a non-blocking manner
	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting %s: %v\n", command, err)
		return
	}

	// Countdown from the specified duration
	for i := duration; i > 0; i-- {
		fmt.Printf("Killing %s in %d seconds...\n", command, i)
		time.Sleep(1 * time.Second)
	}

	// Kill the command
	if err := cmd.Process.Kill(); err != nil {
		fmt.Printf("Error killing %s: %v\n", command, err)
		return
	}

	fmt.Printf("%s was killed\n", command)
}
