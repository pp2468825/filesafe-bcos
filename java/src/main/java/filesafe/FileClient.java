package filesafe;

import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TSocket;
import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.nio.ByteBuffer;
 
public class FileClient
{
    /**
     * 类内测试 main 函数
     * */
    public static void main(String[] args)
    {

        // 检索服务器文件
        SearchFileReq req = new SearchFileReq();
        req.user_id = "13787287171";

        // 上传文件路径
        String filePath = "./filesafe/3.txt";
        // 构造文件数据
        byte[] bytes = toByteArray(filePath);
        AddFileReq add_req = new AddFileReq();
        add_req.file_path = filePath;
        add_req.buff = ByteBuffer.wrap(bytes);
 
        // 构造Thrift客户端，发起请求
        try
        {
            TSocket socket = new TSocket("localhost", 12345);
            socket.setSocketTimeout(60 * 1000);
            TFramedTransport framedTransport = new TFramedTransport(socket);
            framedTransport.open();
            TBinaryProtocol binaryProtocol = new TBinaryProtocol(framedTransport);
            InternalService.Client client = new InternalService.Client(binaryProtocol);
            SearchFileRes rsp = client.search_file(req);
            System.out.println(rsp);
            AddFileRes add_rsp = client.add_file(add_req);
            System.out.println(add_rsp);
        }
        catch (Exception x)
        {
            x.printStackTrace();
        }
    }
 
    /**
     * 文件转化为字节数组
     * */
    private static byte[] toByteArray(String filePath){
        byte[] buffer = null;
        try {
            File file = new File(filePath);
            FileInputStream fis = new FileInputStream(file);
            ByteArrayOutputStream bos = new ByteArrayOutputStream(1000);
            byte[] b = new byte[1000];
            int n;
            while ((n = fis.read(b)) != -1) {
                bos.write(b, 0, n);
            }
            fis.close();
            bos.close();
            buffer = bos.toByteArray();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return buffer;
    }
}
