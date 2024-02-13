package logger

import (
	"dummy-proj-internal/pkg/internal/auth"
	"fmt"
)

func DoLog() {
	fmt.Println("doing some log")
	auth.IsAuth()
}
