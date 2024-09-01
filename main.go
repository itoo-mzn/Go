package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	// エクセルを作成
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Hello")
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
