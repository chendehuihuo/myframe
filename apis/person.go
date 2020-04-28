package apis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	rd "myframe/assets"
	"myframe/models"
	"net/http"
	"strconv"
)

func init (){

}

func  PersonIndex(c *gin.Context) {
	conn:=rd.Redis
	_, err := conn.Do("SET", "name", "go conf set one")
	if err != nil {
		log.Fatalln(err)
	}
	_,err=conn.Do("expire","name",30)
	if err != nil {
		log.Fatalln(err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, name)
}

//插入
func  PersonPost(c *gin.Context) {
	name := c.Request.FormValue("name")
	time := c.Request.FormValue("time")
	//time := c.Param("time")
	times, err := strconv.Atoi(time)

	p := models.Name{Name: name, Time: times}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//获取全部
func  PersonAll(c *gin.Context) {
	var name models.Name
	re, err := name.GetPersons()

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": re,
	})
}

//获取单个
func  PersonOne(c *gin.Context) {
	id := c.Request.FormValue("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}
	p := models.Name{Id: ids}
	ra := p.GetPersonsone()
	c.JSON(http.StatusOK, gin.H{
		"person": ra,
	})
}

func  PersonPut(c *gin.Context) {
	id := c.Request.FormValue("id")
	ids, err := strconv.Atoi(id)
	person := models.Name{Id: ids}

	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := person.PutPersonsone()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update name %d successful %d", person.Name, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func  PersonDelete(c *gin.Context) {
	id := c.Request.FormValue("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}
	person := models.Name{Id: ids}
	ra, err := person.DeletePersonone()

	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
