package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const NoRowFound = "record not found"

func NewInternalServiceErr(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewBadRequestErr(field string) gin.H {
	return gin.H{"error": fmt.Sprintf("%s field is invalid", field)}
}

func NewNotFoundErr(resource string, params interface{}) gin.H {
	return gin.H{"error": fmt.Sprintf("%s(#%v) is not found", resource, params)}
}
