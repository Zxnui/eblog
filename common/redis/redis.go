package model

type RedisManager struct {
}

func (c *RedisManager) Get(key string) string{
	rc:=RedisClient.Get()
	value,err:=rc.Do("get",key)
	if err!=nil{
		return nil
	}
	defer rc.Close()
	return value
}