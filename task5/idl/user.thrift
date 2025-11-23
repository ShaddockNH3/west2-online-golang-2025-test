include "common.thrift" 

namespace go user

struct RegisterUserRequest{
    1: string username
    2: string password
}

struct RegisterUserResponse{
    1: common.BaseResponse base
}

struct LoginUserRequest{
    1: string username
    2: string password
    3: optional string code
}

struct LoginUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct InfoUserRequest{
    1: string user_id
}

struct InfoUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct AvatarUploadUserRequest{
    1: string user_id
    2: binary data
}

struct AvatarUploadUserResponse{
    1: common.BaseResponse base
    2: common.UserDataResponse data
}

struct QrcodeMFAAuthRequest{
    1: string user_id
    // 开启 MFA 时使用
}

struct QrcodeMFAAuthResponse{
    1: common.BaseResponse base
    2: common.QrcodeMFAAuthResponse data
}

struct BindMFAAuthRequest{
    1: string user_id
    2: string code
    3: string secret
}

struct BindMFAAuthResponse{
    1: common.BaseResponse base
}

struct SearchImageRequest{
    1: binary data
}

struct SearchImageResponse{
    1: common.BaseResponse base
    2: string data // 搜索结果，为单个url
}

service UserService {
   RegisterUserResponse RegisterUser(1:RegisterUserRequest req)
   LoginUserResponse LoginUser(1:LoginUserRequest req)
   InfoUserResponse InfoUser(1:InfoUserRequest req)
   AvatarUploadUserResponse AvatarUploadUser(1:AvatarUploadUserRequest req)
   QrcodeMFAAuthResponse QrcodeMFAAuth(1:QrcodeMFAAuthRequest req)
   BindMFAAuthResponse BindMFAAuth(1:BindMFAAuthRequest req)
   SearchImageResponse SearchImage(1:SearchImageRequest req)
}