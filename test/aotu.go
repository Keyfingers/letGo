package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func main() {

	url := "https://www.socialscrm.com/wscrm-bus-api/open/action/batchAdd"
	method := "POST"
	tenantId := "503516"
	userName := "MOBENE"
	apiKey := "571C40CF271D5F9F1C4DB884A58962EA"
	createTime := time.Now().Format("2006-01-02 15:04:05")

	genToken := MD5(tenantId + userName + createTime + apiKey)

	data := "[{\"userName\":\"mobene010\",\"whatsId\":\"8615810078006\",\"friendWhatsId\":\"8617396631093\",\"type\":1}]"
	token := genToken

	post := "{\"tenantId\":\"" + tenantId +
		"\",\"userName\":\"" + userName +
		"\",\"data\":" + data +
		",\"token\":\"" + token +
		"\",\"createTime\":\"" + createTime +
		"\"}"

	fmt.Printf("\n\n%s\n\n", post)

	var jsonStr = []byte(post)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s", body)
}
