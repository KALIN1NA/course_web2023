package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []FeaturedPostData
	MostRecentPosts []MostResentPostData
}

type FeaturedPostData struct {
	Title          string
	Subtitle       string
	Author         string
	Date           string
	ImageAuthor    string
	InfoAuthor     string
}

type MostResentPostData struct {
	Title          string
	Subtitle       string
	Author         string
	Date           string
	ImageAuthor    string
	InfoAuthor     string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: FeaturedPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func FeaturedPosts() []FeaturedPostData {
	return []FeaturedPostData{
		{
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
		    Author: "Mat Vogels",
			Date: "September 25, 2015",
			ImageAuthor: "static/image/Vogels.jpg", 
			InfoAuthor: "Mat Vogels",
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year.",
			Author: "William Wong",
			Date: "September 25, 2015",
			ImageAuthor: "static/image/Wong.jpg",
			InfoAuthor: "William Wong",
		},
	}
}

func MostRecentPosts() []MostResentPostData {
	return []MostResentPostData {
		{
			Title:    "Still Standing Tall",
			Subtitle: "Life begins at the end of your comfort zone.",
		    Author: "William Wong",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Wong.jpg",
			InfoAuthor: "William Wong",
		},
		{
			Title:    "Sunny Side Up",
			Subtitle: "No place is ever as bad as they tell you it’s going to be.",
		    Author: "Mat Vogels",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Vogels.jpg", 
			InfoAuthor: "Mat Vogels",
		},
		{
			Title:    "Water Falls",
			Subtitle: "We travel not to escape life, but for life not to escape us.",
		    Author: "Mat Vogels",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Vogels.jpg", 
			InfoAuthor: "Mat Vogels",
		},
		{
			Title:    "Through the Mist",
			Subtitle: "Travel makes you see what a tiny place you occupy in the world.",
		    Author: "William Wong",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Wong.jpg",
			InfoAuthor: "William Wong",
		},
		{
			Title:    "Awaken Early",
			Subtitle: "Not all those who wander are lost.",
		    Author: "Mat Vogels",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Vogels.jpg", 
			InfoAuthor: "Mat Vogels",
		},
		{
			Title: "Try it Always",
			Subtitle: "The world is a book, and those who do not travel read only one page.",
		    Author: "Mat Vogels",
			Date: "9/25/2015",
			ImageAuthor: "static/image/Vogels.jpg", 
			InfoAuthor: "Mat Vogels",
		},
	}
}


