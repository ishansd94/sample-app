package remote_broker

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"github.com/sanity-io/litter"

	"github.com/ishansd94/sample-app/pkg/log"
)

func Get(c *gin.Context) {
	var server = melody.New()

	server.HandleMessage(func(session *melody.Session, msg []byte) {
		litter.Dump(string(msg))
		fmt.Errorf("")
	})

	err := server.HandleRequest(c.Writer, c.Request)
	if err != nil {
		log.Error("Broker", "Websocket Error", err)
	}
}

func Post(c *gin.Context) {

}
