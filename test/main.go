package main

import (
	"github.com/noneandundefined/vision-go"
	"github.com/noneandundefined/vision-go/config"
	"github.com/noneandundefined/vision-go/pkg/vemail"
)

func main() {
	vision := vision.NewVision()

	// CONFIG
	config.EMAIL_CLIENT = "artemvlasiv1909@gmail.com"
	config.EMAIL_PASSWORD = "bdhdqbfglvdxuqsx"

	vemail.EmailStats(vision)
}
