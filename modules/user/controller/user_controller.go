package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
    service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
    return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    req, err := validation.ValidateCreateUser(c)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := ctrl.service.CreateUser(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context){
    idParam := c.Param("id")
    
    id, err := strconv.Atoi(idParam)
    if err != nil{
        utils.ErrorResponse(c, http.StatusBadRequest, "Format ID tidak valid")
        return
    }
    user, err := ctrl.service.GetUserByID(uint(id))
    if err != nil{
        if errors.Is(err, gorm.ErrRecordNotFound){
            utils.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan")
            return
        }
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data user")
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil data user", user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context){
   users, err := ctrl.service.GetAllUsers()
   if err != nil{
    utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data semua user")
    return
   }
   utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil semua data user", users)
}
