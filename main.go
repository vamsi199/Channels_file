package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

type content struct {
	C1 string `csv:"c1"`
	C2 string `csv:"c2"`
	C3 string `csv:"c3"`
	C4 string `csv:"url"`
	C5 string `csv:"c5"`
	C6 string `csv:"c6"`
	}

func main() {
	var c content


	input,err:=os.Open("input.csv")
	if err!=nil {
		fmt.Errorf("cannot open the file",err)
	}
	r := csv.NewReader(input)
	r.Comma = '\t'
	rec,err:=r.ReadAll()
	var urls []string
	channel := make(chan string,10)
	for _,val := range rec {
		c.C4 = val[3]
		urls = append(urls,c.C4)
		}

	fmt.Println(urls)
	var cols5,cols6 []string

	for _,url:=range urls {
		channel <- url
		c.C5,c.C6=go f(<- channel)
		cols5=append(cols5,c.C5)
		cols6=append(cols6,c.C6)
		}

	fmt.Println(cols5,cols6)



	}
func f(url string)(col4 string, col5 string){
	/*switch url{
	case “xyz11”
		return “d1”, “e1"
		//like this add for all urls xyz2 to xyz4
	}*/
	fmt.Println(url)
	var c1,c2 string
	c1 = "1"
	c2 = "2"
	return c1,c2
}
