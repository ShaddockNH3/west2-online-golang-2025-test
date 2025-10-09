namespace go task3

enum Code{
    Success=1
    ParamInvalid=2
    DBErr=3
}

struct User{
    1: i64 user_id
    2: string name
    3: string introduce
}

struct CreateUserRequest{
    1: string name (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    2: string introduce (api.body="introduce", api.form="introduce",api.vd="(len($) > 0 && len($) < 1000)")
    3: string password (api.body="password", api.form="password", api.vd="(len($) >= 6)")
}

struct CreateUserResponse{
    1: Code Code
    2: string msg
}

struct QueryUserRequest{
    1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
    2: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryUserResponse{
    1: Code Code
    2: string msg
    3: list<User> Users
    4: i64 total
}

struct DeleteUserRequest{
    1: i64 user_id (api.path="user_id",api.vd="$>0")
}

struct DeleteUserResponse{
    1: Code Code
    2: string msg
}

struct UpdateUserRequest{
    1: i64    user_id   (api.path="user_id",api.vd="$>0")
    2: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    3: string introduce (api.body="introduce", api.form="introduce",api.vd="(len($) > 0 && len($) < 1000)")
}

struct UpdateUserResponse{
    1: Code code
    2: string msg
}


// jwt认证相关
struct LoginRequest {
    1: string name (api.body="name", api.form="name", api.vd="(len($) > 0 && len($) < 100)")
    2: string password (api.body="password", api.form="password", api.vd="(len($) > 0)")
}

struct LoginResponse {
    1: Code code
    2: string msg
    3: optional string token
}

service UserService {
   UpdateUserResponse UpdateUser(1:UpdateUserRequest req)(api.post="/v1/user/update/:user_id")
   DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.post="/v1/user/delete/:user_id")
   QueryUserResponse  QueryUser(1: QueryUserRequest req)(api.post="/v1/user/query/")
   CreateUserResponse CreateUser(1:CreateUserRequest req)(api.post="/v1/user/create/")
   LoginResponse Login(1:LoginRequest req)(api.post="/v1/user/login/")
}

// 待办事项
enum Status{
    ToDo=0
    Complete=1
}

// 四个参数
struct ToDoList{
    1: i64 user_id
    2: i64 todo_list_id
    3: Status status
    4: string title
    5: string context
}

struct CreateToDoListRequest{
    1: string title (api.body="title",api.form="title",api.vd="len($)>0&&len($)<1000")
    2: string context (api.body="context",api.form="context",api.vd="len($)>0&&len($)<1000")
}

struct CreateToDoListResponse{
    1: Code code
    2: string msg
}

struct UpdateToDoListRequest{
    1: i64 todo_list_id (api.path="todo_list_id",api.form="todo_list_id",api.vd="$>0")
    2: Status status (api.path="status",api.form="status",api.vd="($==0||$==1)")
    3: string title (api.body="title",api.form="title",api.vd="len($)>0&&len($)<1000")
    4: string context (api.body="context",api.form="context",api.vd="len($)>0&&len($)<1000")
}

struct UpdateToDoListResponse{
    1: Code code
    2: string msg
}

struct QueryToDoListRequest{
    1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
    2: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryToDoListResponse{
    1: Code code
    2: string msg
    3: list<ToDoList> todo_list
    4: i64 total
}

struct DeleteToDoListRequest{
    1: i64 todo_list_id (api.path="todo_list_id",api.vd="$>0")
}

struct DeleteToDoListResponse{
    1: Code code
    2: string msg
}

struct DeleteCompletedToDoListsRequest {
}

struct DeleteCompletedToDoListsResponse {
    1: Code code
    2: string msg
}

struct DeleteAllUserToDoListsRequest {
}

struct DeleteAllUserToDoListsResponse {
    1: Code code
    2: string msg
}

service ToDoListService{
    CreateToDoListResponse CreateToDoList(1:CreateToDoListRequest req)(api.post="v1/todo_list/create/")
    UpdateToDoListResponse UpdateToDoList(1:UpdateToDoListRequest req)(api.post="v1/todo_list/update/:todo_list_id")
    QueryToDoListResponse QueryToDoList(1:QueryToDoListRequest req)(api.post="v1/todo_list/query/")
    DeleteToDoListResponse DeleteToDoList(1:DeleteToDoListRequest req)(api.post="v1/todo_list/delete/:todo_list_id")
    DeleteCompletedToDoListsResponse DeleteCompletedToDoLists(1: DeleteCompletedToDoListsRequest req)(api.post="/v1/todo_list/delete_completed")
    DeleteAllUserToDoListsResponse DeleteAllUserToDoLists(1: DeleteAllUserToDoListsRequest req)(api.post="/v1/todo_list/delete_all")
}