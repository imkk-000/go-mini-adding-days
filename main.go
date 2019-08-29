package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const day = time.Hour * 24
const format = "2006-01-02"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("please put adding days")
		return
	}
	addingDays, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("parsing days err:", err)
		return
	}
	now := time.Now()
	fmt.Println(now.Format(format))
	fmt.Println(now.Add(day * time.Duration(addingDays)).Format(format))
}
