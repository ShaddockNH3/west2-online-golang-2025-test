package redis

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
)

func SetMFASecretToCache(userID string, secret string) error {
	key := userID + constants.MFAAuthSuffix
	return rdbMFAAuth.Set(key, secret, mfaSetupExpireTime).Err()
}

func GetMFASecretFromCache(userID string) (string, error) {
	key := userID + constants.MFAAuthSuffix
	return rdbMFAAuth.Get(key).Result()
}

func DelMFASecretFromCache(userID string) error {
	key := userID + constants.MFAAuthSuffix
	return rdbMFAAuth.Del(key).Err()
}
