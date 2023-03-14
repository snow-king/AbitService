package PersonControllers

import (
	"AbitService/app/service"
	"AbitService/app/service/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Не верный id"}})
	}
	person := new(service.PersonService)
	response := person.Show(id)
	c.JSON(http.StatusOK, response)
}
func ShowFamily(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Не верный id"}})
	}
	person := new(service.PersonService)
	response := person.GetFamily(id)
	c.JSON(http.StatusOK, response)
}
