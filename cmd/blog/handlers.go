package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPageData struct {
	Title           string
	Subtitle        string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postPageData struct{
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Content     string `db:"content"`
	ImageUrl   string `db:"image_url"`
	Author      string `db:"author"`
}

type featuredPostData struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishData string `db:"publish_date"`
	ImageUrlFP  string `db:"image_urlFP"`       //name
}

type mostRecentPostData struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishData string `db:"publish_date"`
	ImageUrlMr      string `db:"image_url"`           //name
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		mostRecentPostsData, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		data := indexPageData {
			Title:  "Let's do it together",
			Subtitle: "We travel the world in search of stories. Come along for the ride.",
      		FeaturedPosts:   featuredPostsData,
			MostRecentPosts: mostRecentPostsData,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] 

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			fmt.Println(postIDStr)
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	
		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

	}
}

func postByID(db *sqlx.DB, postID int) (postPageData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			content,
			image_url,
			author
		FROM 
			post
		WHERE
			post_id = ?
	`

	var post postPageData

	err := db.Get(&post, query, postID)
	if err != nil {
		return postPageData{}, err
	}

	return post, nil
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
		    post_id,
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_urlFP
		FROM
			post
		WHERE featured = 1
	`

	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
		    post_id,
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url
		FROM
			post
		WHERE featured = 0
	`

	var posts []mostRecentPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
