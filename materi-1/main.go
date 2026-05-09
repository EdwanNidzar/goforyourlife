package main

import (
	"fmt"
	"runtime"
)

func main() {
	const appName = "SysInfo Tool"
	version := "1.0.0"

	fmt.Printf("Aplikasi: %s\n", appName)
	fmt.Printf("Versi: %s\n", version)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Go Version: %s\n", runtime.Version())
}