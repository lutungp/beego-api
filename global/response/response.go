package response

const (
	ERROR   = 0
	SUCCESS = 1
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
