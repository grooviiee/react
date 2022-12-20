package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver"
)

// reference -> https://opendart.fss.or.kr/guide/detail.do?apiGrpCd=DS001&apiId=2019001
const apikey string = "1436a1b25f6279f51ed0daf3719a04f1f0f2a333"
const query string = "https://opendart.fss.or.kr/api/list.json?crtfc_key=" + apikey + "&corp_code=aasdd&bgn_de=20200117&end_de=20200117&corp_cls=Y&page_no=1&page_count=10"
const queryCorpCode string = "https://opendart.fss.or.kr/api/corpCode.xml?crtfc_key=" + apikey

//GET	https://opendart.fss.or.kr/api/list.json

type Corp struct {
	List []struct {
		CorpCode   string `xml:"corp_code"`
		CorpName   string `xml:"corp_name"`
		StockCode  string `xml:"stock_code"`
		ModifyDate string `xml:"modify_date"`
	} `xml:"list"`
}

func setupRouter(data string) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, data)
	})
	return r
}

func main() {

	//DownloadCorpCode()

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
	//fmt.Printf("%-v", string(fileReadData))
	OperateService(fmt.Sprintf("%v", corp))

}

func OperateService(data string) error {
	r := setupRouter(data)
	r.Run(":83")

	return nil
}

func DownloadCorpCode() error {
	//Send Query
	resp, err := http.Get(queryCorpCode)
	if err != nil {
		panic(err)
		return err
	}
	defer resp.Body.Close()

	//Get Resp
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return err
	}

	// Make zip file using resp
	err = ioutil.WriteFile("data.zip", data, 0644)

	// Unzip downloaded file
	err = archiver.Unarchive("data.zip", "Unarchive_output")
	if err != nil {
		panic(err)
	}
	err = archiver.Walk("data.zip", func(f archiver.File) error {
		fmt.Println(f.Name(), f.Size())
		return nil
	})
	if err != nil {
		panic(err)
		return err
	}

	return nil
}
