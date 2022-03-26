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

	s.txt += fmt.Sprintf("\n\n %s", time.Now().Format("02 January 03:04"))
	s.txt += fmt.Sprintf("%s", message)

	// Remove some messages when there are too many
	maxLettersPerRow := 44
	maxNrRows := 15
	for {
		nrWrappedLines := len(s.txt) / maxLettersPerRow
		nrNewLines := strings.Count(s.txt, "\n")
		fmt.Println("nr new lines: ", nrNewLines)
		fmt.Println("nr wrapped lines: ", nrWrappedLines)
		if nrNewLines+nrWrappedLines < maxNrRows {
			break
		}
		fmt.Println(s.txt)
		idx := strings.Index(s.txt, "\n")
		if idx == -1 { // Just a single long line
			s.txt = s.txt[2*maxLettersPerRow:]
		} else if idx == 0 {
			s.txt = s.txt[idx+1:]
		} else {
			s.txt = s.txt[idx:]
		}

		// Remove new line from beginning of string
		for {
			if strings.HasPrefix(s.txt, "\n") {
				s.txt = s.txt[1:]
			} else {
				break
			}
		}
	}

	c.Status(http.StatusOK)
}

func (s *Service) httpGETActivityState(c *gin.Context) {
	c.String(http.StatusOK, s.txt)
}
