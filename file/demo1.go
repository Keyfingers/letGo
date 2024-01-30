package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
)

func main() {
	http.HandleFunc("/download-excel", func(w http.ResponseWriter, r *http.Request) {
		f := excelize.NewFile()
		f.SetCellValue("Sheet1", "A1", "mobile")
		f.SetCellValue("Sheet1", "B1", "email")

		//设置内容
		f.SetCellValue("Sheet1", "A2", "+5236781128341")
		f.SetCellValue("Sheet1", "B2", "1232424@163.com")

		fileName := "example.xlsx"

		if err := f.SaveAs(fileName); err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		http.ServeFile(w, r, fileName)
	})

	http.ListenAndServe(":8087", nil)
}
