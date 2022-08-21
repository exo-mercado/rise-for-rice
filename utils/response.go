package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

type ErrorMsg struct {
    Key         string `json:"field"`
    Error     string `json:"message"`
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
    fmt.Println("Im running")
    fmt.Println(err)
    var ve validator.ValidationErrors
    var sqlErr mysql.MySQLError

    if errors.As(err, &ve) {
        fmt.Println("---")
        fmt.Println(ve)
        out := make([]ErrorMsg, len(ve))
        for i, fe := range ve {
            out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
        }
        c.AbortWithStatusJSON(statusCode, gin.H{"errors": out})
    }

    if( errors.Is(err, &sqlErr)) {
            switch sqlErr.Number {
                case 1062:
                    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Duplicate entry"})
                default:
                    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error"})
            }
    }

    // fmt.Println("------")
    // fmt.Println(mysqlErr)
    // fmt.Println("------")
    // fmt.Println("------")

}