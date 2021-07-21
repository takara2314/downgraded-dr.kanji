package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"

	"tcj3-kadai-tuika-kun/services/answer"
	"tcj3-kadai-tuika-kun/services/quiz"
	"tcj3-kadai-tuika-kun/types"
)

var (
	bot      *linebot.Client
	config   types.ConfigYaml
	flexQuiz []byte
)

func init() {
	var err error

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	configData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	quiz.Config = &config
	answer.Config = &config

	flexQuiz, err = ioutil.ReadFile("./templates/quizInfo.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
