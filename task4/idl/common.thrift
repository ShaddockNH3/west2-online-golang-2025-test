namespace go common

// BaseResponse
struct BaseResponse{
    1: string code
    2: string msg
}

// User

struct User{
    1: string id
    2: string username
    3: string password
    4: string avatar_url
    5: string create_at
    6: string update_at
    7: string delete_at
}

struct UserDataResponse{
    1: string id
    2: string username
    3: string avatar_url
    4: string create_at
    5: string update_at
    6: string delete_at
}

struct VideoItems{
    1: string id
    2: string user_id
    3: string video_url
    4: string cover_url
    5: string title
    6: string description
    7: i64 visit_count
    8: i64 like_count
    9: i64 comment_count
    10: string create_at
    11: string update_at
    12: string delete_at
}

struct VideoDataForPopularResponse{
    1: list<VideoItems> items
}

struct VideoDataForListResponse{
    1: list<VideoItems> items
    2: i64 total
}

// interact

struct CommentItems{
    1: string id
    2: string user_id
    3: string video_id
    4: string parent_id
    5: string like_count
    6: string child_count
    7: string content 
    8: string create_at
    9: string update_at
    10: string delete_at
}

struct CommentDataForListResponse{
    1: list<CommentItems> items
}

// social

struct SocialItems{
    1: string id
    2: string user_id
    3: string username
}

struct SocialDataForListResponse{
    1: list<SocialItems> items
    2: i64 total
}