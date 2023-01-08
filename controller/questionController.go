package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quiz-api/models"
)

func Question() gin.HandlerFunc {
	return func(c *gin.Context) {
		var question []models.Question
		category, limit, difficulty := userParams(c)
		if limit == "" {
			limit = "1"
		}

		if difficulty == "" {
			difficulty = "easy"
		}

		url := fmt.Sprintf("https://the-trivia-api.com/api/questions?limit=%s&categories=%s&difficulty=%s", limit, category, difficulty)
		resp, err := http.Get(url)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		defer resp.Body.Close()
		// fmt.Println(resp)
		res, err := io.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		//fmt.Println(string(res))
		if errMar := json.Unmarshal(res, &question); errMar != nil {
			fmt.Println(errMar)
		}

		c.JSON(http.StatusOK, gin.H{"data": question, "length": len(question)})
	}
}

func userParams(c *gin.Context) (string, string, string) {
	category := string(c.Params.ByName("category"))
	limit := string(c.Params.ByName("limit"))
	difficulty := string(c.Params.ByName("difficulty"))

	return category, limit, difficulty
}
