package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// ErrQuestionNotFound is an error raised when a product cannot be found in the database
var ErrQuestionNotFound = fmt.Errorf("Product not found")

// Question defines the structure for an API question
//swagger:model
type Question struct {
	Question string `json:"question"`
	OptionA  string `json:"option_a"`
	OptionB  string `json:"option_b"`
	OptionC  string `json:"option_c"`
	OptionD  string `json:"option_d"`
	Answer   string `json:"answer"`
}

// Questions defines a slice of Question
type Questions []*Question

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// GetQuestions returns all questions from the database
func GetQuestions() Questions {
	return questionsList
}

// GetQuestion returns a single question which matches the id from the database.
// If a question is not found, this function returns a QuestionNotFoundError
func GetQuestion(id int) (*Question, error) {
	for i := range questionsList {
		if i == id {
			return questionsList[id], nil
		}
	}
	return nil, ErrQuestionNotFound
}

var questionsList = func() Questions {
	var qList Questions

	csvfile, err := os.Open("./data/questions.csv")
	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		qList = append(qList, &Question{
			Question: record[1],
			OptionA:  record[2],
			OptionB:  record[3],
			OptionC:  record[4],
			OptionD:  record[5],
			Answer:   record[6],
		})
	}
	return qList
}()
