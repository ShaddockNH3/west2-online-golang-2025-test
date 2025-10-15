package redis

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/go-redis/redis/v7"
)

func AddRdbPopular(k string, v int64) {
	add(rdbPopular, k+constants.PopularVideosSuffix, v)
}

func DelRdbPopular(k string, v int64) {
	del(rdbPopular, k+constants.PopularVideosSuffix, v)
}

func CheckRdbPopular(k string) bool {
	return check(rdbPopular, k+constants.PopularVideosSuffix)
}

func ExistRdbPopular(k string, v int64) bool {
	return exist(rdbPopular, k+constants.PopularVideosSuffix, v)
}

func CountRdbPopular(k string) (sum int64, err error) {
	return count(rdbPopular, k+constants.PopularVideosSuffix)
}

func GetRdbPopular(k string) (vt []int64) {
	get(rdbPopular, k+constants.PopularVideosSuffix)
	return
}

func GetHotVideoIDs(key string, start, end int64) ([]string, error) {
	return getHotIDs(rdbPopular, key, start, end)
}

func getHotIDs(rdb *redis.Client, key string, start, end int64) ([]string, error) {
	result, err := rdb.ZRevRange(key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AddOrUpdatePopularVideo(videoID string, score float64) error {
	// Key: 统一的排行榜名字 constants.PopularVideosSuffix
	// Member: 视频的ID
	// Score: 视频的热度分数
	err := rdbPopular.ZAdd(constants.PopularVideosSuffix, &redis.Z{
		Score:  score,
		Member: videoID,
	}).Err()

	return err
}
