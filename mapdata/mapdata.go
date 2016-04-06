package mapdata

import ()

//Define map structure
type Dobj struct {
	Data  map[string]int
	Queue chan string
}

func NewDobj() Dobj {
	var d Dobj
	d.Data = make(map[string]int)
	d.Queue = make(chan string, 1000)
	go d.writeToMap()
	return d
}

//Define interface to write map
func (d Dobj) Write(word string) {
	go d.writeToChan(word)
}

func (d Dobj) writeToChan(word string) {
	d.Queue <- word
}

func (d Dobj) writeToMap() {
	for {
		w := <-d.Queue
		if v, ok := d.Data[w]; ok {
			d.Data[w] = v + 1
		} else {
			d.Data[w] = 1
		}
	}
}

//Define interface to read from map
func (d Dobj) Read(word string) (e error, v int) {
	return e, d.Data[word]
}

//Define interface to read all from map
