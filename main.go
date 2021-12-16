package main

import (
	"log"
	"regexp"
	"sort"
)

type sorter struct {
	val int
	str string
}

func counter() []sorter {
	var s []sorter
	text := "For those who believe in God, most of the big questions are answered. But for those of us who can't readily accept the God formula, the big answers don't remain stone-written. We adjust to new conditions and discoveries. We are pliable. Love need not be a command nor faith a dictum. I am my own god. We are here to unlearn the teachings of the church, state, and our educational system. We are here to drink beer. We are here to kill war. We are here to laugh at the odds and live our lives so well that Death will tremble to take us."
	arrayText := regexp.MustCompile("[^0-9a-zA-Z]+").Split(text, -1)

	m := map[string]int{}
	for _, word := range arrayText {
		m[word]++
	}
	for k, v := range m {
		s = append(s, sorter{v, k})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].val > s[j].val
	})

	return s[:10]
}

func main() {

	for i, s := range counter() {
		log.Println(i, s.str, s.val)
	}

}
