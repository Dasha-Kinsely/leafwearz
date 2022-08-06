package startup

import (
	"leafwearz/globals"
	"sync"

	"github.com/gin-gonic/gin"
)

func InitServer(wg *sync.WaitGroup) {
	defer wg.Done()
	globals.DefaultServer = gin.Default()
	AddCommonMiddleware()
	return
}