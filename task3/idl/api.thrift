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
    1: optional string Keyword (api.body="keyword")
    2: i64 page (api.body="page",api.query="page",api.vd="$ > 0")
    3: i64 page_size (api.body="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
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
   UpdateUserResponse UpdateUser(1:UpdateUserRequest req)(api.put="/v1/users/:user_id")
   DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.delete="/v1/users/:user_id")
   QueryUserResponse  QueryUser(1: QueryUserRequest req)(api.get="/v1/users")
   CreateUserResponse CreateUser(1:CreateUserRequest req)(api.post="/v1/users")
   LoginResponse Login(1:LoginRequest req)(api.post="/v1/user/login")
}

// 待办事项
enum Status{
    PENDING=0
    COMPLETED=1
    // ARCHIVED=2
}

// 这里还得再改，设计一下todolist的返回格式
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
    1: required i64 todo_list_id (api.path="todo_list_id"),
    2: optional string title (api.body="title"), 
    3: optional string context (api.body="context"),
    4: optional Status status (api.body="status"),
}

struct UpdateToDoListResponse{
    1: Code code
    2: string msg
}

struct UpdateBatchStatusRequest {
    1: required Status status (api.body="status")
}

struct UpdateBatchStatusResponse{
    1: Code code
    2: string msg
}

struct QueryBatchToDoListsRequest{
    1: optional string keyword (api.query="keyword")
    2: optional Status status (api.query="status")
    3: i64 page (api.query="page",api.vd="$ > 0")
    4: i64 page_size (api.query="page_size",api.vd="($ > 0 && $ <= 100)")
}

struct QueryBatchToDoListResponse{
    1: Code code
    2: string msg
    3: list<ToDoList> todo_lists
    4: i64 total
}

struct DeleteToDoListRequest{
    1: i64 todo_list_id (api.path="todo_list_id",api.vd="$>0")
}

struct DeleteToDoListResponse{
    1: Code code
    2: string msg
}

struct DeletePatchToDoListRequest{
}

struct DeletePatchToDoListResponse{
    1: Code code
    2: string msg
}

service ToDoListService{
    CreateToDoListResponse CreateToDoList(1:CreateToDoListRequest req)(api.post="/v1/todo_lists")

    UpdateToDoListResponse UpdateToDoList(1:UpdateToDoListRequest req)(api.patch="/v1/todo_lists/:todo_list_id")
    UpdateBatchStatusResponse UpdateBatchStatus(1:UpdateBatchStatusRequest req)(api.put="/v1/todo_lists/status")

    QueryBatchToDoListResponse QueryBatchToDoList(1:QueryBatchToDoListsRequest req)(api.get="/v1/todo_lists")

    DeleteToDoListResponse DeleteToDoList(1:DeleteToDoListRequest req)(api.delete="/v1/todo_lists/:todo_list_id")

    DeletePatchToDoListResponse DeletePendingToDos(1:DeletePatchToDoListRequest  req)(api.delete="/v1/todo_lists/pending")
    DeletePatchToDoListResponse DeleteCompletedToDos(1:DeletePatchToDoListRequest req)(api.delete="/v1/todo_lists/completed")
    DeletePatchToDoListResponse DeleteAllToDos(1:DeletePatchToDoListRequest req)(api.delete="/v1/todo_lists")
}