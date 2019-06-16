package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
}

var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {
	r := gin.Default()
	// Set static files
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			hostName, err := os.Hostname()
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Pong!",
				"hostName": hostName,
			})
		})
	}
	api.GET("/jokes", jokeHandler)
	api.POST("/jokes/like/:jokeID", likeJoke)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func jokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

func likeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes += 1
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
