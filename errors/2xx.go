package errors

/////////////////////////////////
// 2xx errors variables
/////////////////////////////////
var (
	OK              = &ArgError{ArgErrorSystemMarket, 200, "Операция завершена успешно", "Операция завершена успешно"}
	ContentNotFound = &ArgError{ArgErrorSystemMarket, 204, "content not found", "content not found"}
)
