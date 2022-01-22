package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "github.com/gin-gonic/gin"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "root"
  password = "root"
  dbname   = "driving-license-examination"
)

var config = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

type Question struct {
    Id     int `json:"id"`
    Image  string `json:"image"`
    Detail string `json:"detail"`
}
type Response struct {
  Question Question `json:"question"`
  Choices []Choice `json:"choices"`
}

type Choice struct {
    Id         int `json:"id"`
    QuestionId int `json:"questionId"`
    Image      string `json:"image"`
    Detail     string `json:"detail"`
    IsCorrect  bool   `json:"isCorrect"`
}
//
type QuestionService struct {
  database *sql.DB
}

func (QS QuestionService) getRandomQuestion() (Question, error) {
  query := "SELECT * FROM questions ORDER BY random() LIMIT 1"

  var id int
	var image, detail string
  err := QS.database.QueryRow(query).Scan(&id, &image, &detail);
  if err != nil {
    return Question{}, err
  }
  return Question{
    Id: id,
    Image: image,
    Detail: detail,
  }, nil 
}

type ChoiceService struct {
  database *sql.DB
}

func (CS ChoiceService) getChoicesByQuestionId(questionId int) ([]Choice, error) {
  rows, err := CS.database.Query("SELECT * FROM choices WHERE question_id = $1", questionId)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var choices []Choice

  for rows.Next() {
    var choice Choice
    // when scan we need to that these orders align order with database
    if err := rows.Scan(
      &choice.Id, &choice.Image, &choice.Detail,
      &choice.IsCorrect, &choice.QuestionId); err != nil {
      return choices, err
    }
    choices = append(choices, choice)
  }
  if err = rows.Err(); err != nil {
    return choices, err
  }
  return choices, nil
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func CheckDB() {
  // open database
  db, err := sql.Open("postgres", config)
  CheckError(err);
  defer db.Close()
  err = db.Ping()
  CheckError(err)
  fmt.Println("Connected!")
}

func main() {
  db, err := sql.Open("postgres", config)
  CheckError(err);
  defer db.Close()
  
  
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware


	router := gin.Default()

	router.GET("/questions/random", func(c *gin.Context){
    question, err := QuestionService{db}.getRandomQuestion();
    if err != nil {
      panic(err)
    }
    choices, err := ChoiceService{db}.getChoicesByQuestionId(question.Id);
    if err != nil {
      panic(err)
    }
    c.JSON(200, Response{
      Question:question,
      Choices:choices,
    })
  })
	router.Run()
}

