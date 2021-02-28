package redis

import (
	"errors"
	"strconv"
	"time"
)

func (c *redisClient) SaveAuth(signedToken string, userID uint) error {
	err := c.client.Set(signedToken, userID, time.Second*86400).Err()
	return err
}

func (c *redisClient) GetAuth(signedToken string) (uint, error) {
	strID, err := c.client.Get(signedToken).Result()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(strID)
	return uint(userID), err
}

func (c *redisClient) RemoveAuth(signedToken string) error {
	err := c.client.Del(signedToken)
	if err == nil {
		return errors.New("Can not delete key")
	}
	return nil
}
