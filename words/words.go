package words

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"word_stat/mapdata"
)

var start, end int64

func WordsFromPara(para string, d mapdata.Dobj) {
	bytes := []byte(para)
	p1 := 0
	p2 := 0
	for k, v := range bytes {
		if v < 39 || (v > 39 && v < 65) || v > 122 || (v > 90 && v < 97) {
			p2 = k
			if p2-p1 > 1 {
				//word = Start P1 and End p2
				word := string(bytes[p1:p2])
				d.Write(word)
			}
			p1 = p2 + 1
		}
	}
}

func WordsFromFile(filename string, d mapdata.Dobj) {
	cd, _ := os.Getwd()
	file := cd + "/" + filename
	start = time.Now().UnixNano()
	if f, err := os.Open(file); err == nil {
		if _, err = f.Seek(0, 0); err == nil {
			buf := bufio.NewReader(f)
			for {
				line, err := buf.ReadString('\n')
				if err == io.EOF {
					break
				} else if len(line) > 1 {
					go WordsFromPara(line, d)
				}
			}
			end = time.Now().UnixNano()
			fmt.Println("Para spliting done and took", (end-start)/1000000, "MS")
		} else {
			fmt.Println("Fail to seek")
		}
	} else {
		fmt.Println("Fail to open file", file)
	}
}
