package resp

type Response struct {
	Success bool        `cbt:"success"`
	Data    interface{} `cbt:"data"`
	Err     FailData    `cbt:"err"`
}

type FailData struct {
	ErrorCode string
	ErrorMsg  string
}
