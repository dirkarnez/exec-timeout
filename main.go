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

// 	if err != nil {
// 		fmt.Println("If the command was killed, err will be \"signal: killed\"")
// 		fmt.Println("If the command wasn't killed, it contains the actual error, e.g. invalid command")
// 	} else {
// 		fmt.Println("Done")
// 	}
// }
package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Starting taskmgr...")

	// Start taskmgr in a non-blocking manner
	cmd := exec.Command("taskmgr")
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting taskmgr:", err)
		return
	}

	// Countdown from 5 seconds
	for i := 5; i > 0; i-- {
		fmt.Printf("Killing taskmgr in %d seconds...\n", i)
		time.Sleep(1 * time.Second)
	}

	// Kill taskmgr
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("Error killing taskmgr:", err)
		return
	}

	fmt.Println("taskmgr was killed")
}
