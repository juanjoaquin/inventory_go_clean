package dtos

/* DTOs son Data Transfer Objects. Son objetos que se usan para transferir datos entre el cliente y el servidor. */
/* Tambien se aplicaran las validaciones de los datos en los DTOs. */
type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}
