package users

// **UserHandler struct**
type UserHandler struct {
	service UserService
}

// **NewUserHandler Constructor**
func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}
