package common

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v3"
)

const (
	Version    = "4Q"
	ServiceURL = "https://downgraded-drkanji.appspot.com"
)

var (
	Bot          *linebot.Client
	VisionAPICtx context.Context
	VisionAPI    *vision.ImageAnnotatorClient

	Quizzes       QuizzesYaml
	AntonymFormat []byte
	HomonymFormat []byte
	SynonymFormat []byte
	ConferFormat  []byte
	WritingFormat []byte
	ReadingFormat []byte

	QuizTypes = []string{
		"Antonym",
		"Homonym",
		"Synonym",
		"Confer",
		"Writing",
		"Reading",
	}

	QuizSections = []string{
		"antonyms",
		"homonyms",
		"synonyms",
		"confers",
		"others",
	}
)

type Quiz struct {
	Type     string
	Section  string
	No       int
	Option   string
	Content  string
	Corrects []string
	Memo     string
}

type QuizzesYaml struct {
	Antonyms [][]string `yaml:"antonyms"`
	Homonyms [][]string `yaml:"homonyms"`
	Synonyms [][]string `yaml:"synonyms"`
	Confers  [][]string `yaml:"confers"`
	Others   []string   `yaml:"others"`
}

func init() {
	var err error

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	// Load a quiz file.
	file, err := ioutil.ReadFile("./quizzes.yaml")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// File data to a instance.
	err = yaml.Unmarshal(file, &Quizzes)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Load quiz templates.
	AntonymFormat, err = ioutil.ReadFile("./templates/Antonym.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	HomonymFormat, err = ioutil.ReadFile("./templates/Homonym.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	SynonymFormat, err = ioutil.ReadFile("./templates/Synonym.json")
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

	// Load and authorize Google Vision API
	VisionAPICtx = context.Background()

	VisionAPI, err = vision.NewImageAnnotatorClient(VisionAPICtx)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
