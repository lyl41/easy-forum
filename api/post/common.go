package post

type Post struct {
}

type HttpReply struct {
	Err error  `json:"err"`
	Msg string `json:"msg"`
}