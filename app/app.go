package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func init()  {
	fmt.Println(`
                             _                                    _               
                            (_)                                  (_)              
  __ _  ___ ______ _ __ ___  _  ___ _ __ ___  ___  ___ _ ____   ___  ___ ___  ___ 
 / _  |/ _ \______| '_   _ \| |/ __| '__/ _ \/ __|/ _ \ '__\ \ / / |/ __/ _ \/ __|
| (_| | (_) |     | | | | | | | (__| | | (_) \__ \  __/ |   \ V /| | (_|  __/\__ \
 \__, |\___/      |_| |_| |_|_|\___|_|  \___/|___/\___|_|    \_/ |_|\___\___||___/
  __/ |                                                                           
 |___/                                                                            

	`)
}

// StartApp function to initialize app
func StartApp() {
	MapUrls()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
