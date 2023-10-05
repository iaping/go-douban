package main

import (
	"fmt"

	"github.com/iaping/go-douban/movie"
)

func main() {
	sub, err := movie.New().LoadSubject("https://movie.douban.com/subject/36576596/", map[string]string{})
	if err != nil {
		panic(err)
	}
	fmt.Println(sub.Poster())
	fmt.Println(sub.Title())
	fmt.Println(sub.Summary())
	fmt.Println(sub.Average())
	fmt.Println(sub.Country())
	fmt.Println(sub.Language())
	fmt.Println(sub.Release())
	fmt.Println(sub.Genre())
	fmt.Println(sub.Celebrities())
}
