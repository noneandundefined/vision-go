package main

import (
	"github.com/noneandundefined/vision-go"
	"github.com/noneandundefined/vision-go/pkg/vemail"
	"github.com/noneandundefined/vision-go/vconfig"
)

func main() {
	vision := vision.NewVision()

	// CONFIG
	vconfig.EMAIL_CLIENT = "artemvlasiv1909@gmail.com"
	vconfig.ATTACH_LOGFILES = true
	vconfig.LOGFILES_BY_TIME_STYLES = true

	vemail.EmailStats(vision)
}
