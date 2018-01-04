package wells

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Wells(station string)(wells []string){
	f, err := os.Open(station+".txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := bufio.NewReader(f)
	defer f.Close()

	var list []string

	for {
		line, err := buf.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		list = append(list, strings.Replace(string(line), "\r\n", "", -1))
	}
	return list
}