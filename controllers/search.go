package controllers

import (
	"strings"
)

type Top struct {
	Topic string
}

func AllTitles() []string {
	data := Decoded()
	var Titles []string
	for _, d := range data {
		Titles = append(Titles, d.Title)
	}
	return Titles
}

func TopTen() []Top {
	data := Decoded()
	var Topdata []Top
	for i, top := range data {
		if i == 10 {
			break
		}
		topten := Top{
			Topic: top.Topic,
		}
		Topdata = append(Topdata, topten)
	}
	return Topdata
}

// func uniqueTitles(intSlice []int) []int {
// 	keys := make(map[int]bool)
// 	list := []int{}
// 	for _, entry := range intSlice {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			list = append(list, entry)
// 		}
// 	}
// 	return list
// }

func Search(keyword string) []string {
	titles := AllTitles()
	// log.Println(len(titles))
	var indexes []int
	for i := 0; i < len(titles); i++ {
		exists := strings.Contains(titles[i], keyword)
		if exists == true {
			indexes = append(indexes, i)
		}
		// title := len(titles[i])
		// for j := 0; j < title; j++ {
		// 	// newj := j
		// 	// fmt.Println("j", j)
		// 	// again := len(keyword) + j
		// 	// fmt.Println("again", again)
		// 	// matching := string(titles[i][newj:again])
		// 	matching := string(titles[i][j])
		// 	if keyword == matching {
		// 		// fmt.Println(i)
		//
		// 	}
		// 	// j += again
		// }
	}
	// fmt.Println(indexes)
	// NotRepeatIndex := uniqueTitles(indexes)
	// // fmt.Println(NotRepeatIndex)
	var Title []string
	for _, num := range indexes {
		// fmt.Println(titles[num])
		Title = append(Title, titles[num])
		// fmt.Println("-----------------------------------------------------------------------------------------")
	}
	return Title
}
