package serializer

// Response 基础序列化器
type Response struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  string      `json:"error"`
}


//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

