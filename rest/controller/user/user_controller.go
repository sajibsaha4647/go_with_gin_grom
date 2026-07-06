package user

import (
	"ecommerce/domain"
	"ecommerce/dto"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	// 1. Validate email format
	if !utils.IsValidEmail(req.Email) {
		utils.SendError(c, 400, "invalid email address")
		return
	}

	// 2. Check if user exists
	user, err := h.repo.FindByEmail(req.Email)
	if err != nil {
		utils.SendError(c, 404, "Email not found")
		return
	}

	if user == nil {
		utils.SendError(c, 400, "Invalid email or password")
		return
	}

	// 3. Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.SendError(c, 400, "Invalid email or password")
		return
	}

	user.Password = ""

	utils.SendSuccess(c, 200, "Login successful", user)
}

func (h *UserHandler) register(c *gin.Context) {
	var req dto.RegisterRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	// 1. Validate email format
	if !utils.IsValidEmail(req.Email) {
		utils.SendError(c, 400, "invalid email address")
		return
	}

	// 3. Check duplicate email
	exists, err := h.repo.ExistsByEmail(req.Email)
	if err != nil {
		utils.SendError(c, 500, "Internal server error")
		return
	}

	if exists {
		utils.SendError(c, 400, "User with this email already exists")
		return
	}

	if len(req.Password) < 8 {
		utils.SendError(c, 400, "Password must be at least 8 characters")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.SendError(c, 500, "Internal server error")
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		UserType: "customer",
	}

	// Save to database
	if err := h.repo.Store(user); err != nil {
		utils.SendError(c, 500, "Failed to create user")
		return
	}

	user.Password = ""

	utils.SendSuccess(c, 201, "User registered successfully", user)

}

func (h *UserHandler) singleUser(c *gin.Context) {

	idStr := c.Param("id")

	// id, err := strconv.ParseUint(idStr, 10, 32)
	// if err != nil {
	// 	utils.SendError(c, 400, "Invalid user id")
	// 	return
	// }

	user, err := h.repo.SingleUser(idStr)
	if err != nil {
		utils.SendError(c, 404, "User not found")
		return
	}

	utils.SendSuccess(c, 200, "User found", user)
}

func (h *UserHandler) userlist(c *gin.Context) {}

func (h *UserHandler) updateUser(c *gin.Context) {}

func (h *UserHandler) deleteUser(c *gin.Context) {}

func (h *UserHandler) rowCount(c *gin.Context) {}
