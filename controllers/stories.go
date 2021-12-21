package controllers

import (
	"github.com/gin-gonic/gin"
	"mall/database"
)

type Story struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"body"`
}

func Show(c * gin.Context) {
	db := database.DBConn()

	rows, err := db.Query("SELECT * FROM stories LIMIT 100")
	if err != nil {
		c.JSON(500, gin.H {
			"message": err.Error(),
		})
	}

	var listStories []Story

	for rows.Next() {
		var id int
		var title, body string
		story := Story{}

		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		story.Id = id
		story.Title = title
		story.Content = body

		listStories = append(listStories, story)
	}

	c.JSON(200, listStories)
	defer db.Close()
}

func Create(c * gin.Context) {
	db := database.DBConn()

	type CreateStory struct {
		Title string `form:"title" json: "title" binding:"required"`
		Body string `form:"body" json: "body" binding:"required"`
	}

	var json CreateStory

	if err := c.ShouldBindJSON(&json); err == nil {
		insStory, err := db.Prepare("INSERT INTO stories(title, body) VALUE(?,?)")
		if err != nil {
			c.JSON(500, gin.H {
				"message": err,
			})
		}

		insStory.Exec(json.Title, json.Body)
		c.JSON(200, gin.H {
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	}

	defer db.Close()
}

func Read(c * gin.Context) {
	db := database.DBConn()
	rows, err := db.Query("SELECT id, title, body FROM stories WHERE id = " + c.Param("id") + " LIMIT 1")
	if err != nil {
		c.JSON(500, gin.H {
			"message": "Story not found",
		})
	}

	story := Story{}

	for rows.Next() {
		var id int
		var title, body string

		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		story.Id = id
		story.Title = title
		story.Content = body
	}

	c.JSON(200, story)
	defer db.Close()
}

func Update(c * gin.Context) {
	db := database.DBConn()
	type UpdateStory struct {
		Title string `form:"title json:"title" binding:"required"`
		Body string `form:"body" json:"body" binding:"required"`
	}

	var json UpdateStory
	if err := c.ShouldBindJSON(&json); err == nil {
		edit, err := db.Prepare("UPDATE stories SET title = ?, body = ? WHERE id = " + c.Param("id"))
		if err != nil {
			panic(err.Error())
		}
		edit.Exec(json.Title, json.Body)

		c.JSON(200, gin.H {
			"message": "edited",
		})
	} else {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	}
	defer db.Close()
}

func Delete(c * gin.Context) {
	db := database.DBConn()
	delete, err := db.Prepare("DELETE FROM stories WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(c.Param("id"))
	c.JSON(200, gin.H {
		"message": "deleted",
	})

	defer db.Close()
}