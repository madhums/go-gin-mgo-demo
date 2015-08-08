package articles

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/madhums/go-gin-mgo-demo/models"
)

// New article
func New(c *gin.Context) {
	article := models.Article{}

	c.HTML(http.StatusOK, "articles/form", gin.H{
		"title":   "New article",
		"article": article,
	})
}

// Create an article
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionArticle).Insert(article)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}

// Edit an article
func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionArticle).FindId(oID).One(&article)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "articles/form", gin.H{
		"title":   "Edit article",
		"article": article,
	})
}

// List all articles
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	articles := []models.Article{}
	err := db.C(models.CollectionArticle).Find(nil).Sort("-updated_on").All(&articles)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "articles/list", gin.H{
		"title":    "Articles",
		"articles": articles,
	})
}

// Update an article
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"title":      article.Title,
		"body":       article.Body,
		"updated_on": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionArticle).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}

// Delete an article
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionArticle).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}
