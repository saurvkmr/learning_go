package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	images := make(chan string)
	maps := make(chan string)
	videos := make(chan string)
	news := make(chan string)

	var search []string

	go getImages(images)
	go getMaps(maps)
	go getVideos(videos)
	go getNews(news)

	for {
		select {
		case image, ok := <-images:
			search = append(search, image)
			if !ok {
				images = nil
			}

		case mapsres, ok := <-maps:
			search = append(search, mapsres)
			if !ok {
				maps = nil
			}
		case video, ok := <-videos:
			search = append(search, video)
			if !ok {
				videos = nil
			}
		case newsres, ok := <-news:
			search = append(search, newsres)
			if !ok {
				news = nil
			}
		}

		if images == nil && maps == nil && videos == nil && news == nil {
			break
		}
	}

	fmt.Println(search)

}

func getImages(images chan string) {
	for i := 0; i < 10; i++ {
		images <- ("image" + strconv.Itoa(i))
		time.Sleep(10 * time.Millisecond)
	}
	close(images)
}

func getMaps(maps chan string) {
	for i := 0; i < 10; i++ {
		maps <- ("maps" + strconv.Itoa(i))
		time.Sleep(10 * time.Millisecond)
	}
	close(maps)
}
func getVideos(videos chan string) {
	for i := 0; i < 10; i++ {
		videos <- ("videos" + strconv.Itoa(i))
		time.Sleep(10 * time.Millisecond)
	}
	close(videos)
}
func getNews(news chan string) {
	for i := 0; i < 10; i++ {
		news <- ("news" + strconv.Itoa(i))
		time.Sleep(10 * time.Millisecond)
	}
	close(news)
}
