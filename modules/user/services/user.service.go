package services

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/modules/user/dtos"
	"belajar-go-fiber/modules/user/models"
	"belajar-go-fiber/utils"
)

// UserService provides user-related business logic

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers(page, pageSize int, sortBy, sortOrder, keyword string) ([]*dtos.UserResponse, int, int, error) {
	db := database.DB.Model(&models.User{})
	var users []models.User
	var totalData int64

	// Set default values if not provided
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if sortBy == "" {
		sortBy = "id"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	offset := (page - 1) * pageSize
	order := sortBy + " " + sortOrder

	if keyword != "" {
		db = db.Where(
			"COALESCE(name, '') LIKE ? OR COALESCE(email, '') LIKE ? OR COALESCE(phone, '') LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%",
		)
	}

	// Count total data
	if err := db.Where("deleted_at IS NULL").Count(&totalData).Error; err != nil {
		return nil, 0, 0, err
	}

	if err := db.Order(order).Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, 0, err
	}

	var userResponses []*dtos.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &dtos.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		})
	}
	totalPage := int((totalData + int64(pageSize) - 1) / int64(pageSize)) // ceil division
	return userResponses, int(totalData), totalPage, nil
}

func (s *UserService) GetUserById(id int) (*dtos.UserResponse, error) {
	var user models.User
	if err := database.DB.Debug().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &dtos.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Phone: user.Phone}, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(userReq *dtos.UserCreateRequest) (*models.User, error) {
	newUser := models.User{
		Name:  userReq.Name,
		Email: userReq.Email,
		Phone: userReq.Phone,
	}

	hashPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashPassword

	if err := database.DB.Debug().Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (s *UserService) UpdateUser(userReq *dtos.UserUpdateRequest) (*models.User, error) {
	updatedUser := models.User{
		ID:    userReq.ID,
		Name:  userReq.Name,
		Email: userReq.Email,
		Phone: userReq.Phone,
	}

	if userReq.Password != nil {
		hashPassword, err := utils.HashPassword(*userReq.Password)
		if err != nil {
			return nil, err
		}
		updatedUser.Password = hashPassword
	}

	if err := database.DB.Debug().Model(&models.User{}).Where("id = ?", userReq.ID).Updates(updatedUser).Error; err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (s *UserService) DeleteUser(id int) error {
	if err := database.DB.Debug().Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
