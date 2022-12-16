package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

// reference -> https://opendart.fss.or.kr/guide/detail.do?apiGrpCd=DS001&apiId=2019001
const apikey string = "1436a1b25f6279f51ed0daf3719a04f1f0f2a333"
const query string = "https://opendart.fss.or.kr/api/list.json?crtfc_key=" + apikey + "&corp_code=aasdd&bgn_de=20200117&end_de=20200117&corp_cls=Y&page_no=1&page_count=10"
const queryCorpCode string = "https://opendart.fss.or.kr/api/corpCode.xml?crtfc_key=" + apikey

//GET	https://opendart.fss.or.kr/api/list.json

type Corp struct {
	corp_code   string
	corp_name   string
	stock_code  string
	modify_date string
}

func setupRouter(data string) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {

	// resp, err := http.Get(queryCorpCode)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// err = ioutil.WriteFile("data.zip", data, 0644)

	// err = archiver.Unarchive("data.zip", "Unarchive_output")
	// if err != nil {
	// 	panic(err)
	// }
	// err = archiver.Walk("data.zip", func(f archiver.File) error {
	// 	fmt.Println(f.Name(), f.Size())
	// 	return nil
	// })
	// if err != nil {
	// 	panic(err)
	// }

	fp, err := os.Open("Unarchive_output\\CORPCODE.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	fileReadData, _ := ioutil.ReadAll(fp)

	// xml 디코딩
	var corp []Corp
	xmlerr := xml.Unmarshal(fileReadData, &corp)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Printf("%-v", corp)

	r := setupRouter(string("ping"))
	r.Run(":8080")
}
