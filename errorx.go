package errorx

const defaultCode = -1

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Ok(msg string, data ...any) error {
	var Data any
	if len(data) > 0 {
		Data = data[0]
	} else {
		Data = nil
	}
	return &CodeError{Code: 0, Msg: msg, Data: Data}
}
func NewCode(code int, msg any, data ...any) error {
	var Data any
	if len(data) > 0 {
		Data = data[0]
	} else {
		Data = nil
	}
	switch e := msg.(type) {
	case error:
		return &CodeError{Code: code, Msg: e.Error(), Data: Data}
	case string:
		return &CodeError{Code: code, Msg: e, Data: Data}
	default:
		return &CodeError{Code: code, Msg: "never", Data: Data}
	}
}

func New(msg any, data ...any) error {
	var Data any
	if len(data) > 0 {
		Data = data[0]
	} else {
		Data = nil
	}
	return NewCode(defaultCode, msg, Data)
}

func (e *CodeError) Error() string {
	return e.Msg
}
