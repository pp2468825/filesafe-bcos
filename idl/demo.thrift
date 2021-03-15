
namespace java service.demo

struct FileDataReq
{
    1:required string   name,                   // 文件名字
    2:required binary   buff,                   // 文件数据
}

struct FileDataRsp
{
    1:required i32 code,
    2:required string msg,
}

service FileInfoExtractService
{
    FileDataRsp uploadFile(1:FileDataReq filedata);       // 文件解析函数
}