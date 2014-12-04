package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dir := flag.String("path", "db/migrations/", "goose migrations path")
	seed := flag.String("seed", "", "Init count date as string (20060102150405)")
	flag.Parse()

	var serie Serie
	var err error
	var init_date time.Time

	if strings.EqualFold(*seed, "") == false {
		if init_date, err = time.Parse("2006/01/02 15:04 05", *seed); err != nil {
			log.Printf("Seed must be in date time format as 2006/01/02 15:04 05 \n %v", err)
			return
		}
	} else {
		init_date = time.Now()
	}
	if serie, err = NewSerie(init_date); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	list, _ := ioutil.ReadDir(*dir)
	for _, file := range list {
		if new_name, err := UpdateName(file, *dir, &serie); err != nil {
			log.Println("Error: %v", err)
		} else {
			log.Printf("File %v renamed to %v \n", file.Name(), new_name)
		}
	}
}

func UpdateName(file os.FileInfo, path string, serie *Serie) (new_name string, err error) {
	oldname := file.Name()
	name := strings.SplitAfter(oldname, "_")
	serie.Next()

	oldname = fmt.Sprintf("%v/%v", path, oldname)
	filename := fmt.Sprintf("%v/%v_%v", path, serie.Formated, name[1])
	return filename, os.Rename(oldname, filename)
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
