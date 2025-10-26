package redis

import (
	"fmt"
	"strconv"

	goRedis "github.com/go-redis/redis/v7"
)

const (
	KeyLikeSet   = "likes:%s:%s"        // e.g. likes:video:video_id
	KeyLikeCount = "counts:%s_likes:%s" // e.g. counts:video_likes:video_id
)

func AddLike(userID, objectType, objectID string) (bool, error) {
	key := fmt.Sprintf(KeyLikeSet, objectType, objectID)
	added, err := rdbLike.SAdd(key, userID).Result()
	if err != nil {
		return false, err
	}
	// 如果是新点赞，则增加点赞数
	if added == 1 {
		countKey := fmt.Sprintf(KeyLikeCount, objectType, objectID)
		if err = rdbLike.Incr(countKey).Err(); err != nil {
			// 如果增加点赞数失败，撤销刚才的点赞操作
			rdbLike.SRem(key, userID)
			return false, err
		}
		return true, nil
	}
	// 用户已经点过赞
	return false, nil
}

func RemoveLike(userID, objectType, objectID string) (bool, error) {
	key := fmt.Sprintf(KeyLikeSet, objectType, objectID)
	removed, err := rdbLike.SRem(key, userID).Result()
	if err != nil {
		return false, err
	}
	// 如果是取消点赞，则减少点赞数
	if removed == 1 {
		countKey := fmt.Sprintf(KeyLikeCount, objectType, objectID)
		if err = rdbLike.Decr(countKey).Err(); err != nil {
			// 如果减少点赞数失败，撤销刚才的取消点赞操作
			rdbLike.SAdd(key, userID)
			return false, err
		}
		return true, nil
	}
	// 用户没有点过赞
	return false, nil
}

func GetLikeCount(objectType, objectID string) (int64, error) {
	countKey := fmt.Sprintf(KeyLikeCount, objectType, objectID)
	countStr, err := rdbLike.Get(countKey).Result()
	if err != nil {
		// 说明是0个赞
		if err == goRedis.Nil {
			return 0, nil
		}
		return 0, err
	}
	return strconv.ParseInt(countStr, 10, 64)
}

func IsLiked(userID, objectType, objectID string) (bool, error) {
	key := fmt.Sprintf(KeyLikeSet, objectType, objectID)
	return rdbLike.SIsMember(key, userID).Result()
}
