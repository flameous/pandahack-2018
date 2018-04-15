package server

import (
	"github.com/flameous/pandahack-2018/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"strconv"
	"log"
)

type server struct {
	db   *db.Database
	port string
}

func NewServer(db *db.Database, port string) *server {
	return &server{db, port}
}

func (s *server) Serve() {
	engine := gin.New()

	engine.Use(cors.Default())
	engine.Use(gin.Logger())

	engine.GET("user/:id", s.getTaskHandler(s.db.GetUserByID))
	engine.GET("task/:id", s.getTaskHandler(s.db.GetTaskByID))
	engine.Run(":" + s.port)
}

func (s *server) getTaskHandler(getSomething db.GetterByID) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.Param("id")
		id, err := validateID(raw)
		if err != nil {
			log.Println(raw, err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		task, err := getSomething(id)
		if err != nil {
			log.Println(id, err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, task)
	}
}

func validateID(raw string) (int, error) {
	id, err := strconv.ParseInt(raw, 10, 32)
	return int(id), err
}
