package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success[T any](data T) Response[T] {
	return Response[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Error(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
