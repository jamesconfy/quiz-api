package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quiz-api/models"
)

// Question route godoc
// @Summary	Get questions
// @Description	Question route
// @Tags	Question
// @Accept	json
// @Produce	json
// @Param	category	query	string	false	"category"
// @Param	limit	query	string	false	"limit"
// @Param	difficulty	query	string	false	"difficulty"
// @Success	200  {object}  []models.Question
// @Router	/ [get]
func Question() gin.HandlerFunc {
	return func(c *gin.Context) {
		var question []models.Question
		category, limit, difficulty := userParams(c)
		if limit == "" {
			limit = "10"
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
		// res, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		// }

		//fmt.Println(string(res))
		if errDec := json.NewDecoder(resp.Body).Decode(&question); errDec != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		// if errMar := json.Unmarshal(res, &question); errMar != nil {
		// 	fmt.Println(errMar)
		// }

		c.JSON(http.StatusOK, gin.H{"data": question, "length": len(question), "status": http.StatusOK})
	}
}

func userParams(c *gin.Context) (string, string, string) {
	category := string(c.Query("category"))
	limit := string(c.Query("limit"))
	difficulty := string(c.Query("difficulty"))

	return category, limit, difficulty
}
