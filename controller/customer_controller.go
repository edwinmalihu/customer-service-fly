package controller

import (
	"customer-service/request"
	"customer-service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	LoginUser(*gin.Context)
}

type customerController struct {
	// Di sini Anda bisa menambahkan dependensi nanti, misal:
	// db *gorm.DB
}

func NewCustomerController() CustomerController {
	return &customerController{}
}


func (c customerController) LoginUser(ctx *gin.Context) {
	var req request.Login

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	

	res := response.LoginResponse{
		Username: req.Username,
		Msg:      "Login Success",
	}

	ctx.JSON(http.StatusOK, res)
}
