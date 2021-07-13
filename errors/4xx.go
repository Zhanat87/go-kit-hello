package errors

/////////////////////////////////
// 4xx errors variables
/////////////////////////////////
var (
	DynamicCategoryError  = &ArgError{ArgErrorSystemMarket, 400, "Действие допустимо только для статичных рубрик", "Действие недопустимо с динамической рубрикой"}
	ActionNotExist        = &ArgError{ArgErrorSystemMarket, 400, "Запрошено неизвестное действие", "Запрошенного действия не существует"}
	DepthLimitReached     = &ArgError{ArgErrorSystemMarket, 400, "Превышен уровень вложенности", "Превышена 2-х уровневая вложенность категорий"}
	ParentNotExist        = &ArgError{ArgErrorSystemMarket, 400, "Родительская категория не существует", "Попытка наследования несуществующего родителя"}
	InvalidCharacter      = &ArgError{ArgErrorSystemMarket, 400, "incorrect input", "incorrect input"}
	FieldsValidationError = &ArgError{ArgErrorSystemMarket, 400, "Ошибка валидации поля(-ей)", "Некорректные данные"}
	InvalidOTP            = &ArgError{ArgErrorSystemMarket, 400, "Передан неверный OTP", "Передан неверный OTP"}
	Unauthorized          = &ArgError{ArgErrorSystemMarket, 401, "Невалидный авторизационный токен", "Запрос не может быть выполнен из-за конфликтного обращения к ресурсу"}
	AccessDenied          = &ArgError{ArgErrorSystemMarket, 403, "access denied", "access denied"}
	NotFound              = &ArgError{ArgErrorSystemMarket, 404, "not found", "not found"}
	Conflict              = &ArgError{ArgErrorSystemMarket, 409, "resource conflict", "resource conflict"}
	CassandraSaveError    = &ArgError{ArgErrorSystemMarket, 409, "cassandra write error", "cassandra write error"}
	CassandraReadError    = &ArgError{ArgErrorSystemMarket, 409, "cassandra read error", "cassandra read error"}
	CassandraDeleteError  = &ArgError{ArgErrorSystemMarket, 409, "cassandra delete error", "cassandra delete error"}
	CsvError              = &ArgError{ArgErrorSystemMarket, 409, "Данные входного файла Неправильны", "Неправильный JSON"}
	DeserializeBug        = &ArgError{ArgErrorSystemMarket, 415, "de/serialization bug", "de/serialization bug"}
)
