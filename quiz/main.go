package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type QuizResult struct {
	Questions []QuizQuestion
}

type QuizQuestion struct {
	Question string
	Answer string
	Answered string
}

func (q QuizQuestion) String() string {
	return fmt.Sprintf("Question : %s  -  Correct answer : %s  -  Your answer : %s", q.Question, q.Answer, q.Answered)
}

func (r QuizResult) String() string {
	var res string

	for i, q := range r.Questions {
		res += fmt.Sprintf("%d  -  %s\n", i + 1, q)
	}
	res += fmt.Sprintf("Result : Good answers = %d  - Wrong answers = %d\n", r.Correct(), r.Wrong())

	return res
}

func (q *QuizResult) Correct() int {
	var tot int

	for _, question := range q.Questions {
		if question.Answer == question.Answered {
			tot++
		}
	}

	return tot
}

func (q *QuizResult) Wrong() int {
	var tot int

	for _, question := range q.Questions {
		if question.Answer != question.Answered {
			tot++
		}
	}

	return tot
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	result := new(QuizResult)
	buf, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Fatal("Quiz file does not exists !")
	}

	r := csv.NewReader(bytes.NewBuffer(buf))

	for  {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s : ", rec[0])
		a, _ := reader.ReadString('\n')
		result.Questions = append(result.Questions, QuizQuestion{Question: rec[0], Answer: rec[1], Answered: strings.TrimSpace(a)})
	}

	fmt.Print(result)
}