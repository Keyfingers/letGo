package main

import (
	"fmt"
	"strings"
)

func main() {

	body := "{\\\"account_id\\\":230926830100002100,\\\"order_id\\\":230926830200002348,\\\"invalid_time\\\":1695784680578,\\\"sys_code\\\":\\\"Super\\\"}"

	body = string([]byte(strings.Trim(strings.Replace(string(body), "\\", "", -1), "\"")))
	fmt.Println(body)

}
