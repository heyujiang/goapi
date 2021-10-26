package util

import (
	"github.com/spf13/viper"
	"goapi/config"
	"log"
	"testing"
)

func TestPdf(t *testing.T) {
	//init config
	if err := config.Init(""); err != nil {
		log.Fatal(err.Error())
	}
	htmlPath := viper.GetString("pdf.html_path")
	pdfPath := viper.GetString("pdf.pdf_path")
	Pdf(htmlPath+"lagou.html", pdfPath+"logou.pdf")
}
