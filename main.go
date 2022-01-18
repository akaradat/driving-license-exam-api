package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type Question struct {
  Id uint64
  Image string
  Detail string
}

type QuestionResponse struct {
  Id uint64
  Image string
  Detail string
  Choices []Choice
}

type Choice struct {
  Id uint64
  QuestionId uint64
  Image string
  Detail string
  IsCorrect bool
}

type QuestionService struct {
  database Database
}

func Initialize(database Database) *QuestionService {
  return &QuestionService{database}
}

func (QS QuestionService) getRandomQuestion() {
  question, err := DB.query();
  if err != nil {
    return fmt.Errorf("failed to create temp dir: %v", err)
  }
  return question
}



func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/questions/random", func(c *gin.Context){
    const question = QuestionService.getRandomQuestion()
    c.JSON(200, question)
  })
	router.Run()
}

