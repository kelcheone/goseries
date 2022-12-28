package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {

	f := excelize.NewFile()

	index := f.NewSheet("Sheet2")

	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	f.SetActiveSheet(index)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}

	read_file()

}

func read_file() {
	f, err := excelize.OpenFile("Book1.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		log.Fatal(err)
	}

	println(cell)

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		for _, colCell := range row {
			println(colCell)
		}
	}

}
