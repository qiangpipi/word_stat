package words

import (
	"word_stat/mapdata"
)

func WordFromPara(para string, d mapdata.Dobj) {
	words := []string{"test", "aaa"}
	for _, v := range words {
		d.Write(v)
	}
}

func WordFromFile(File string) {
}
