package htmltopdf

import (
	"github.com/spf13/viper"
	"goapi/config"
	"log"
	"testing"
)

func TestCreatePdf(t *testing.T) {
	//init config
	if err := config.Init(""); err != nil {
		log.Fatal(err.Error())
	}
	htmlPath := viper.GetString("pdf.html_path")
	pdfPath := viper.GetString("pdf.pdf_path")
	CreatePdf(htmlPath+"22 讲通关 Go 语言.html", pdfPath+"22 讲通关 Go 语言.pdf")
}
