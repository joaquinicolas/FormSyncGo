package request;

import (
	"net/http"
	"log"
	"strconv"
	"fmt"
	"io/ioutil"
)

type Response struct {
	StatusCode int
	Status string
	Body []byte
}

func MakeRequest(isFirstTime bool,url string) *Response{
	client := &http.Client{}
	req,err := http.NewRequest("GET",url,nil)
	fmt.Println()
	CheckErr(err)
	req.Header.Add("isFirstTime",strconv.FormatBool(isFirstTime))
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	var body_response []byte
	if resp.StatusCode == 200 {
		body_response,_ = ioutil.ReadAll(resp.Body)
	}else{
		body_response = nil
	}

	return &Response{
		Status:resp.Status,
		StatusCode:resp.StatusCode,
		Body:body_response,
	}
}



func CheckErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}