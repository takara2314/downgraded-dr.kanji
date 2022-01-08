package common

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v3"
)

const VERSION = "4Q"

var (
	Bot      *linebot.Client
	Quizzes  QuizzesYaml
	FlexQuiz []byte
)

type QuizzesYaml struct {
	Antonyms [][]string `yaml:"antonyms"`
	Homonym  [][]string `yaml:"homonym"`
	Synonyms [][]string `yaml:"synonyms"`
	Confer   [][]string `yaml:"confer"`
	Three    [][]string `yaml:"three"`
	Four     [][]string `yaml:"four"`
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
}
