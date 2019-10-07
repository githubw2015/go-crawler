package engine

type Request struct {
	Url        string                   //请求的URL
	ParserFunc func([]byte) ParseResult //处理函数
}
type ParseResult struct {
	Requests []Request     //所有的抓取URL的请求
	Items    []interface{} //切片
}
