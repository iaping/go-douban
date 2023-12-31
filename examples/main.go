package main

import (
	"fmt"

	"github.com/iaping/go-douban/movie"
)

func main() {
	sub, err := movie.New().LoadSubject("https://movie.douban.com/subject/1401118/", map[string]string{})
	if err != nil {
		panic(err)
	}
	fmt.Println(sub.Poster())
	fmt.Println(sub.Name())
	fmt.Println(sub.Type())
	fmt.Println(sub.Summary())
	fmt.Println(sub.Average())
	fmt.Println(sub.Country())
	fmt.Println(sub.Language())
	fmt.Println(sub.Release())
	fmt.Println(sub.Genres())
	fmt.Println(sub.Celebrities())
	fmt.Println(sub.OtherNames())
	fmt.Println(sub.Runtime())
	fmt.Println(sub.Imdb())
	fmt.Println(sub.Episode())
	fmt.Println(sub.EpisodeRuntime())

	fmt.Println("===================================================")

	cele, err := movie.New().LoadCelebrity("https://movie.douban.com/celebrity/1035649/", map[string]string{})
	if err != nil {
		panic(err)
	}
	fmt.Println(cele.Photo())
	fmt.Println(cele.Name())
	fmt.Println(cele.FullName())
	fmt.Println(cele.Sex())
	fmt.Println(cele.Constellation())
	fmt.Println(cele.Birthday())
	fmt.Println(cele.Birthplace())
	fmt.Println(cele.EnglishName())
	fmt.Println(cele.ChineseName())
	fmt.Println(cele.Introduction())
	fmt.Println(cele.Career())
	fmt.Println(cele.Imdb())
}
