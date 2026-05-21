package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  string
}

var users = []User{
	{Name: "Pranav", Age: "19"},
	{Name: "abc", Age: "29"},
}

func main() {

	router := gin.Default()

	router.GET("/get", get)
	router.POST("/post", post)
	router.PUT("/put", put)
	router.DELETE("/delete", delete)

	router.Run()

}

func get(c *gin.Context) {
	message := ""
	for _, user := range users {
		message += "Name: " + user.Name + " Age: " + user.Age + "\n"
	}
	c.String(http.StatusOK, message)
}

func post(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	user := User{Name: name, Age: age}
	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func delete(c *gin.Context) {
	name := c.PostForm("name")
	targetIndex := -1
	for i, user := range users {
		if user.Name == name {
			targetIndex = i
		}
	}
	if targetIndex != -1 {
		users = slices.Delete(users, targetIndex, targetIndex+1)
	}
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func put(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	targetIndex := -1
	for i, user := range users {
		if user.Name == name {
			targetIndex = i
		}
	}
	users[targetIndex].Age = age
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}
