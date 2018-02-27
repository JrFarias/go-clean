package common

// Prefix at router
const Prefix = "v1"

//Error struct
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorMessage function to add default error messages
func ErrorMessage() map[string]Error {
	m := make(map[string]Error)

	m["NoCustomer"] = Error{
		Message: "Usuários não encontrados",
		Code:    1,
	}

	m["CustomerNotFound"] = Error{
		Message: "Usuário não encontrado",
		Code:    2,
	}

	m["CustomerAlreadyExists"] = Error{
		Message: "usuario já existente",
		Code:    3,
	}

	return m
}
