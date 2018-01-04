package xml2txt

import (
	"encoding/xml"
	"strings"
	"os"
	"io"
)

type workbook struct {
	Styles string `xml:"Styles"`
	Worksheet worksheet `xml:"Worksheet"`
}

type worksheet struct {
	Table table `xml:"Table"`
}

type table struct {
	Column []string `xml:"Column"`
	Row []row `xml:"Row"`
}

type row struct {
	Cell []string `xml:"Cell>Data"`
}


func Xml2txt(data []byte)  {
	var result workbook
	xml.Unmarshal(data, &result)
	f2, _ := os.OpenFile("data.txt", os.O_APPEND, 0666)
	b := result.Worksheet.Table.Row
	for i:=3;i<=len(b)-1;i++ {
		io.WriteString(f2, b[i].Cell[0]+","+strings.Replace(b[i].Cell[1], "-","",-1)+","+b[i].Cell[9]+","+b[i].Cell[16]+","+b[i].Cell[17]+"\r\n")
	}
	f2.Close()
}

