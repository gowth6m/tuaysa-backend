package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tuaysa.com/pkg/response"
)

type UserHandler struct {
	Repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User object to be created"
// @Success 201 {object} UserResponse "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/create [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var reqPayload CreateUserRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoUser, err := ConvertCreateUserRequestToUser(reqPayload)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := h.Repo.CreateUser(c.Request.Context(), mongoUser)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Created user successfully", createdUser)
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []UserResponse "Users retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/all [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Repo.GetAllUsers(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Retrieved users successfully", users)
}
