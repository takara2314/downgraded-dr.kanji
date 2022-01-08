package common

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v3"
)

const (
	Version = "4Q"
)

var (
	Bot      *linebot.Client
	Quizzes  QuizzesYaml
	FlexQuiz []byte

	AntonymsFormat []byte
	HomonymFormat  []byte
	SynonymsFormat []byte
	ConferFormat   []byte
	WritingFormat  []byte
	ReadingFormat  []byte

	QuizTypes = []string{
		"Antonyms",
		"Homonym",
		"Synonyms",
		"Confer",
		"Writing",
		"Reading",
	}
)

type Quiz struct {
	Type    string
	No      int
	Option  string
	Content string
	Correct []string
}

type QuizzesYaml struct {
	Antonyms [][]string `yaml:"antonyms"`
	Homonym  [][]string `yaml:"homonym"`
	Synonyms [][]string `yaml:"synonyms"`
	Confer   [][]string `yaml:"confer"`
	Others   []string   `yaml:"others"`
}

func init() {
	var err error

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	file, err := ioutil.ReadFile("./quizzes.yaml")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal(file, &Quizzes)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	FlexQuiz, err = ioutil.ReadFile("./templates/quiz.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	AntonymsFormat, err = ioutil.ReadFile("./templates/Antonyms.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	HomonymFormat, err = ioutil.ReadFile("./templates/Homonym.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	SynonymsFormat, err = ioutil.ReadFile("./templates/Synonyms.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	ConferFormat, err = ioutil.ReadFile("./templates/Confer.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	WritingFormat, err = ioutil.ReadFile("./templates/Writing.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	ReadingFormat, err = ioutil.ReadFile("./templates/Reading.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
