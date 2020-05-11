package apis

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"index/index.html",gin.H{"title":"golang views","header":"this is header","footer":"this is footer"})
}
func Layout(c *gin.Context)  {
	template.ParseFiles("views/layout.html", "views/index/index,html")
	c.HTML(http.StatusOK,"layout.html",gin.H{"title":"golang views"})
}
