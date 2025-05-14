package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func Success[T any](data T) Response[T] {
	return Response[T]{
		Code:    0,
		Data:    data,
		Message: "success",
	}
}

func Error(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Data:    nil,
		Message: message,
	}
}
