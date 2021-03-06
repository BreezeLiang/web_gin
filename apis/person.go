package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	. "web_gin/models"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func ModPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}

	p.GetPerson()
	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.ModPerson()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	p := Person{Id: id}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	person := Person{Id: id}

	person.GetPerson()
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
	fmt.Println(person)

}

func GetPersonsApi(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}

func UpPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))

	person := Person{Id: id}
	err = person.GetPerson() //检查是否存在

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person Not Found!",
		})
	} else {
		person.FirstName = firstName
		person.LastName = lastName

		err = person.UpPerson()
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusOK, gin.H{
				"data": nil,
				"msg":  "Update fail",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": "Person exit!",
				"msg":  "Update Successful!",
			})
		}
	}
}
