package errors

var (
	ErrHandshakeOriginNotAllowed       = NewWebsocketError("request origin is not allowed")
	ErrHandshakeMethodNotAllowed       = NewWebsocketHandshakeError("http request method is not GET")
	ErrHandshakeTokenUpgradeNotFound   = NewWebsocketHandshakeError("token 'upgrade' is not found in http header")
	ErrHandshakeTokenWebsocketNotFound = NewWebsocketHandshakeError("token 'websocket' is not found in http header")
	ErrHandShakeTokenSecNotFound       = NewWebsocketHandshakeError("token 'Sec-Websocket-Key' is not found in http header")
	ErrFuncNotSupportWebsocketVersion  = NewWebsocketFuncNotSupportError("'Sec-Websocket-Version:13' is not found in http header")
)

type WebsocketError struct {
	ErrMsg string
}

func NewWebsocketError(message string) WebsocketError {
	return WebsocketError{ErrMsg: message}
}

func (w WebsocketError) error() string {
	return "websocket error: "
}

func (w WebsocketError) Error() string {
	return w.error() + w.ErrMsg
}

type WebsocketHandshakeError struct {
	WebsocketError
}

func NewWebsocketHandshakeError(message string) WebsocketHandshakeError {
	return WebsocketHandshakeError{WebsocketError{ErrMsg: message}}
}

func (w WebsocketHandshakeError) Error() string {
	return w.WebsocketError.error() + "not a websocket handshake: " + w.ErrMsg
}

type WebsocketFuncNotSupportError struct {
	WebsocketError
}

func NewWebsocketFuncNotSupportError(message string) WebsocketFuncNotSupportError {
	return WebsocketFuncNotSupportError{WebsocketError{ErrMsg: message}}
}

func (w WebsocketFuncNotSupportError) Error() string {
	return w.WebsocketError.error() + "not supported: " + w.ErrMsg
}