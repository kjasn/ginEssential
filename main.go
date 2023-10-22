package main

import (
	"Kjasn/ginEssential/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()

	r = CollectRouter(r)

	err := r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}

}
