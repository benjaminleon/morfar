package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const BindAddr = ":13001"

type Service struct {
	txt string
}

func main() {
	service := Service{
		txt: "Väntar på meddelanden",
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.PUT("/", service.httpPUTActivityState)
	router.GET("/", service.httpGETActivityState)

	err := router.Run(BindAddr)
	if err != nil {
		panic(fmt.Sprintf("Failed to run server. Reason: '%v'", err))
	}
}

func (s *Service) httpPUTActivityState(c *gin.Context) {
	message, queryExists := c.GetQuery("s")

	if !queryExists {
		c.String(http.StatusBadRequest, "Missing either 's' query parameter. E.g. send a PUT request with Postman to 127.0.0.1:13001/?s=hejsan")
		return
	}

	// Remove some messages when there are too many
	for {
		fmt.Println("length of txt is ", len(s.txt))
		if len(s.txt) < 50 {
			break
		}
		fmt.Println(s.txt)
		idx := strings.Index(s.txt, "\n")
		fmt.Println("idx", idx)
		if idx != -1 {
			s.txt = s.txt[idx+1:]
			fmt.Println("New length of txt is ", len(s.txt))
		}
	}

	s.txt += fmt.Sprintf("\n%s", message)

	c.Status(http.StatusOK)
}

func (s *Service) httpGETActivityState(c *gin.Context) {
	c.String(http.StatusOK, s.txt)
}
