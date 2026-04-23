package user

import (
    "github.com/gin-gonic/gin"
    "github.com/Mobilizes/materi-be-alpro/modules/user/controller"
    authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
)

func RegisterUserRoutes(r *gin.RouterGroup, ctrl *controller.UserController, jwtSvc *authService.JWTService) {
    users := r.Group("/users")
    {
        users.POST("", ctrl.CreateUser)

        users.GET("/:id", ctrl.GetUserByID)

        users.GET("", ctrl.GetAllUsers)
    }
}
