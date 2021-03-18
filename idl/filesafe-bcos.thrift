namespace go filesafe
namespace java filesafe

typedef string JsonString

enum StatusCode{
    SUCCESS=0,
    EXCEPTION=1,
}

struct SearchFileReq{
    1: string user_id,                              // 取hash过的手机号
}

struct SearchFileRes{
    1: StatusCode status_code,             // 状态码
    2: string message,                              // 返回信息

    10: JsonString file_list,                       // 文件列表
}

struct CheckFileReq{
    1: string user_id,                               //session_id ?
    
    10: string file_path,                            // 本地文件路径
    20: string digest,                               // 本地文件摘要
}

struct CheckFileRes{
    1: StatusCode status_code,
    2: string message,
}

struct AddFileReq{
    1: string user_id,

    10: string file_path,
    20: string digest,
    30: binary buff
}

struct AddFileRes{
    1: StatusCode status_code,
    2: string message, 
}

struct MoveFileReq{
    1: string user_id,

    10: string file_path,                           // 本地文件路径
    20: string target_path,                         // 目的文件路径
    30: string digest,
}

struct MoveFileRes{
    1: StatusCode status_code,
    2: string message, 
}

struct DeleteFileReq{
    1: string user_id,

    10: string file_path,
    20: string digest,
}

struct DeleteFileRes{
    1: StatusCode status_code,
    2: string message, 
}

struct RecoveryFileReq{
    1: string user_id,

    10: string file_path,
    20: string digest,
}

struct RecoveryFileRes{
    1: StatusCode status_code,
    2: string message, 

    10: binary buff
}


service InternalService{
    // 检索用户文件
    SearchFileRes search_file(1: SearchFileReq req),
    // 校验文件
    CheckFileRes check_file(1: CheckFileReq req),
    // 上传文件
    AddFileRes add_file(1: AddFileReq req),
    // 移动文件
    MoveFileRes move_file(1: MoveFileReq req),
    // 删除文件
    DeleteFileRes delete_file(1: DeleteFileReq req),
    // 恢复文件
    RecoveryFileRes recover_file(1: RecoveryFileReq req),
}