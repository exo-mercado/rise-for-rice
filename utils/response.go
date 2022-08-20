package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {

    switch fe.Tag() {
        case "required":
            return "This field is required"
        case "lte":
            return "Should be less than " + fe.Param()
        case "gte":
            return "Should be greater than " + fe.Param()
        case "email":
            return "Must be in email format"
        default:
            return "Invalid value"
   }
}

func ErrorJSON(c *gin.Context, statusCode int, err error) {

    var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]ErrorMsg, len(ve))
            for i, fe := range ve {
                out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
            }
            c.JSON(statusCode, gin.H{"errors": out})
        }

        mysqlErr := err.(*mysql.MySQLError)
        switch mysqlErr.Number {
            case 1062:
                c.JSON(http.StatusBadRequest, gin.H{"errors": "Duplicate entry"})
            default:
                c.JSON(statusCode, gin.H{"errors": err.Error()})
        }
}