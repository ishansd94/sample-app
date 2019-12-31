package sample

import (
	"net/http"

	k8slib "github.com/ericchiang/k8s"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation"
	gouuid "github.com/satori/go.uuid"

	"github.com/ishansd94/sample-app/internal/pkg/k8s"
	"github.com/ishansd94/sample-app/internal/pkg/response"
	"github.com/ishansd94/sample-app/pkg/log"
)

type Request struct {
	Namespace string            `json:"namespace"`
	Name      string            `json:"name"`
	Content   map[string]string `json:"content"`
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Namespace, validation.Required),
	)
}

func Create(c *gin.Context) {

	var req Request
	if err := c.BindJSON(&req); err != nil {
		log.Error("sample.Create", "error while binding request", err)
		response.Default(c, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		response.Custom(c, http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	if req.Content == nil {
		req.Content = map[string]string{
			"uuid": uuid(),
		}
	}

	client, err := k8sclient()
	if err != nil {
		response.Custom(c, http.StatusInternalServerError, gin.H{"errors": err})
		log.Error("sample.Create", "error while creating k8s client", err)
		return
	}

	h := k8s.NewHandler(client)

	if err := h.CreateSecret(req.Name, req.Namespace, req.Content); err != nil{

		if handleK8sError(c, err) {
			return
		}

		response.Custom(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Default(c, http.StatusCreated)
}


func Get(c *gin.Context){

	ns := c.DefaultQuery("namespace", "")
	name := c.DefaultQuery("name", "")

	if empty(ns) {
		response.Default(c, http.StatusBadRequest)
		return
	}

	client, err := k8sclient()
	if err != nil {
		response.Custom(c, http.StatusInternalServerError, gin.H{"errors": err})
		log.Error("sample.Create", "error while creating k8s client", err)
		return
	}

	h := k8s.NewHandler(client)

	if empty(name) {
		secrets, err := h.AllSecrets(ns)
		if err != nil {
			response.Custom(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response.Custom(c, http.StatusOK, gin.H{"items": secrets.GetItems()})
		return
	}

	secret, err := h.GetSecret(name, ns)
	if err != nil {

		if handleK8sError(c, err) {
			return
		}

		response.Custom(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Custom(c, http.StatusOK, gin.H{"item": secret})
}

func uuid() string {
	return gouuid.NewV4().String()
}

func handleK8sError(c *gin.Context,err error) bool {

	if apierr, ok := err.(*k8slib.APIError); ok {
		if apierr.Code == http.StatusNotFound {
			response.Custom(c, http.StatusNotFound, gin.H{"error": err.Error()})
			return true
		}

		if apierr.Code == http.StatusConflict {
			response.Custom(c, http.StatusConflict, gin.H{"error": err.Error()})
			return true
		}
	}

	return false
}

func k8sclient() (*k8slib.Client, error) {
	client, err := k8slib.NewInClusterClient()
	if err != nil {
		log.Error("sample.k8sclient", "error creating k8 client", err)
		return nil, err
	}

	return client, nil
}

func empty(str string) bool {
	return str == ""
}
