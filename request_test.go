package request

import (
	"testing"
	"math/rand"
	"fmt"
)

func TestGetForms(t *testing.T) {
	resp := MakeRequest((rand.Intn(100) < 90),"http://ncrapps.com/apiwf/api/exchange/sp230015@Ncr.com")
	if resp.StatusCode == 200  {
		fmt.Println(string(resp.Body))
	}
	if resp.StatusCode == 204 {
		fmt.Println("No Content")
	}
	if resp.StatusCode == 500 {
		fmt.Println(resp.Status)
	}
	t.Error()
}
