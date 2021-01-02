package data

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Question struct {
	Question string `json:"question"`
	OptionA  string `json:"optionA"`
	OptionB  string `json:"optionB"`
	OptionC  string `json:"optionC"`
	OptionD  string `json:"optionD"`
	Answer   string `json:"answer"`
}

type Questions []*Question

func (q *Questions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}

func (q *Question) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}

func GetQuestions() Questions {
	return questionsList
}

func GetQuestion(id int) *Question {
	return questionsList[id]
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
