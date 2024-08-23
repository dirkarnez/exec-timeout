package main

import (
    "context"
	"fmt"
	"os"
    "os/exec"
	"time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cmd := exec.CommandContext(ctx, os.Args[0], os.Args[1:]...)

    _, err := cmd.CombinedOutput()

    if (ctx.Err() == context.DeadlineExceeded) {
		fmt.Println("Command was killed")
    }

    if err != nil {
        fmt.Println("If the command was killed, err will be \"signal: killed\"")
		fmt.Println("If the command wasn't killed, it contains the actual error, e.g. invalid command")
    }
}