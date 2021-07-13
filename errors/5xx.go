package errors

/////////////////////////////////
// 5xx errors variables
/////////////////////////////////
var (
	APIConnectError        = &ArgError{ArgErrorSystemMarket, 503, "remote API is unreachable", "remote API is unreachable"}
	CassandraConnectError  = &ArgError{ArgErrorSystemMarket, 503, "cassandra is unreachable", "cassandra is unreachable"}
	ElasticConnectError    = &ArgError{ArgErrorSystemMarket, 503, "elasticSearch is unreachable", "elasticSearch is unreachable"}
	RedisConnectError      = &ArgError{ArgErrorSystemMarket, 503, "redis is unreachable", "redis is unreachable"}
	KafkaProducerError     = &ArgError{ArgErrorSystemMarket, 503, "kafka producer init failed", "kafka producer init failed"}
	KafkaConsumerError     = &ArgError{ArgErrorSystemMarket, 503, "kafka consumer init failed", "kafka consumer init failed"}
	RabbitMQConnectError   = &ArgError{ArgErrorSystemMarket, 503, "rabbitMQ is unreachable", "rabbitMQ is unreachable"}
	RabbitMQPublisherError = &ArgError{ArgErrorSystemMarket, 503, "rabbitMQ publisher init failed", "rabbitMQ publisher init failed"}
	RabbitMQServerError    = &ArgError{ArgErrorSystemMarket, 503, "rabbitMQ server init failed", "rabbitMQ server init failed"}
	S3ConnectError         = &ArgError{ArgErrorSystemMarket, 503, "Сервис недоступен", "Недоступен сервис Ceph"}
	SendOTPError           = &ArgError{ArgErrorSystemMarket, 503, "Ошибка отправки OTP", "Сервис отправки OTP не доступен или вернул ошибку"}
	CheckOTPError          = &ArgError{ArgErrorSystemMarket, 503, "Ошибка проверки OTP", "Сервис проверки OTP не доступен или вернул ошибку"}
	SendLoanRequestError   = &ArgError{ArgErrorSystemMarket, 503, "Ошибка отправки кредитной заявки", "Сервис обработки кредитных заявок не доступен или вернул ошибку"}
	NetworkTimeout         = &ArgError{ArgErrorSystemMarket, 503, "Соединение сброшено по таймауту", "Удаленный узел не ответил в установленный таймаут"}
	BufferError            = &ArgError{ArgErrorSystemMarket, 500, "Ошибка чтения буфера", "Ошибка чтения буфера"}
)
