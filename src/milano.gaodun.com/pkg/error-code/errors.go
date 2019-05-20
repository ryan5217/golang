package error_code

type Error interface {
	error
	GetCode() int
}

type StatusError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (e StatusError) Error() string {
	return e.Message
}

func (e StatusError) GetCode() int {
	return e.Code
}

func New(code int, message string) error {
	return StatusError{
		code,
		message,
		nil,
	}
}
