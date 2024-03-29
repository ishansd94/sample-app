package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation"

	"github.com/ishansd94/sample-app/internal/app/sample/proto"
	"github.com/ishansd94/sample-app/internal/pkg/response"
	"github.com/ishansd94/sample-app/pkg/log"
)

type Request struct {
	Field1 string            `json:"field1"`
	Field2 map[string]string `json:"field2"`
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Field1, validation.Required),
		validation.Field(&r.Field1, validation.Required),
	)
}

// Create godoc
// @Summary Sample Create Endpoint .
// @Description Sample Create Endpoint.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func Create(c *gin.Context) {

	var req Request
	if err := c.BindJSON(&req); err != nil {
		log.Debug("sample.Create", "error while binding request, err: "+err.Error())
		response.Default(c, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		response.Custom(c, http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	response.Default(c, http.StatusCreated)
}

// Get godoc
// @Summary Sample Get Endpoint .
// @Description Sample Get Endpoint.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func Get(c *gin.Context) {
	response.Custom(c, http.StatusOK, gin.H{"item": Request{}})
}

// Proto godoc
// @Summary Sample Protocol Buff Endpoint .
// @Description Sample Protocol Buff Endpoint.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func Proto(c *gin.Context) {

	foo := "foo"
	bar := "bar"

	data := &proto.Sample{
		Foo: foo,
		Bar: bar,
	}
	c.ProtoBuf(http.StatusOK, data)
}