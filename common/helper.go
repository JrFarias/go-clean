package common

// Prefix at router
func Prefix() string {
	return "v1"
}

//Error struct
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

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
