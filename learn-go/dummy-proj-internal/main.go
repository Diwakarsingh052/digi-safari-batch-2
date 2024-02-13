package main

import (
	"dummy-proj-internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

// go mod vendor //

//	go mod tidy
//download deps, required by the go.mod
// it also updates the go.mod // indirect and direct deps
// remove any deps not being used
// run above command after downloading or cloning any project from the internet

// go get github.com/bytedance/sonic@v1.9.1 // go get a specific version
//

func main() {
	logger.DoLog()
	gin.New()

}
