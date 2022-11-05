package errorx

const defaultCode = -1
const SkipError = 10086996

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCode(code int, msg any, data any) error {
	switch e := msg.(type) {
	case error:
		return &CodeError{Code: code, Msg: e.Error(), Data: data}
	case string:
		return &CodeError{Code: code, Msg: e, Data: data}
	default:
		return &CodeError{Code: code, Msg: "never", Data: data}
	}
}

func New(msg any) error {
	return NewCode(defaultCode, msg, nil)
}

func (e *CodeError) Error() string {
	return e.Msg
}
