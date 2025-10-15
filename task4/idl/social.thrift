include "common.thrift" 

namespace go social

enum ActionRelationType{
    FOLLOW=0
    UNFOLLOW=1
}

struct SocialResponse{
    1: common.BaseResponse base
    2: common.SocialDataForListResponse data
}

// 关注操作
struct ActionRelationRequest{
    1: optional string to_user_id (api.form="user_id", api.vd="(len($)==0 || len($) > 0 && len($) < 100)")
    2: optional ActionRelationType action_type (api.form="action_type", api.vd="$ in [0,1]")
}

struct ActionRelationResponse{
    1: common.BaseResponse base
}

// 关注列表

struct ListFollowingRequest{
    1: string user_id (api.query="user_id", api.vd="(len($) > 0 && len($) < 100)")
    2: optional i64 page_num (api.query="page_num", api.vd="( $ > 0 )")
    3: optional i64 page_size (api.query="page_size", api.vd="( $ > 0 && $ < 100 )")
}

// 粉丝列表

struct ListFollowerRequest{
    1: string user_id (api.query="user_id", api.vd="(len($) > 0 && len($) < 100)")
    2: optional i64 page_num (api.query="page_num", api.vd="( $ > 0 )")
    3: optional i64 page_size (api.query="page_size", api.vd="( $ > 0 && $ < 100 )")
}

// 好友列表

struct ListFriendsRequest{
    1: optional i64 page_num (api.query="page_num", api.vd="( $ > 0 )")
    2: optional i64 page_size (api.query="page_size", api.vd="( $ > 0 && $ < 100 )")
}

service SocialService {
   ActionRelationResponse ActionRelation(1:ActionRelationRequest req)(api.post="/v1/relation/action")
   SocialResponse ListFollowing(1:ListFollowingRequest req)(api.get="/v1/following/list")
   SocialResponse ListFollower(1:ListFollowerRequest req)(api.get="/v1/follower/list")
   SocialResponse ListFriends(1:ListFriendsRequest req)(api.get="/v1/friends/list")
}