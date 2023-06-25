package services

type IRedisService interface {
	AddKey(key string, value string) error
	ReadKey(key string) (*string, error)
}
