package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(out), "=")
	fmt.Printf("Temperature %s\n", parts[1])
}
