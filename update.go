package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dir := flag.String("path", "db/migrations/", "goose migrations path")
	flag.Parse()

	var serie Serie
	var err error

	if serie, err = NewSerie(time.Now()); err != nil {
		fmt.Printf("Error: %v", err)
	}

	list, _ := ioutil.ReadDir(*dir)
	for _, file := range list {
		if err := UpdateName(file, *dir, &serie); err != nil {
			fmt.Println("Error: %v", err)
		} else {
			fmt.Printf("File %v renamed", file.Name())
		}
	}
}

func UpdateName(file os.FileInfo, path string, serie *Serie) (err error) {
	oldname := file.Name()
	name := strings.SplitAfter(oldname, "_")
	serie.Next()

	oldname = fmt.Sprintf("%v/%v", path, oldname)
	filename := fmt.Sprintf("%v/%v_%v", path, serie.Formated, name[1])
	return os.Rename(oldname, filename)
}

func NewSerie(init_time time.Time) (serie Serie, err error) {
	serie.Timestamp = init_time
	serie.Formated = fmt.Sprintf("%v", init_time.Format("20060102150405"))
	serie.Value, err = strconv.ParseInt(serie.Formated, 10, 64)
	return serie, err
}

type Serie struct {
	Timestamp time.Time
	Formated  string
	Value     int64
}

func (self *Serie) Next() {
	self.Value = self.Value + 3
	self.Formated = fmt.Sprintf("%v", self.Value)
}
