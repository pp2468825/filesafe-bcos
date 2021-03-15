package service.demo;

import org.apache.thrift.TException;
import java.io.FileOutputStream;
import java.nio.channels.FileChannel;
 
public class FileServiceImpl implements FileInfoExtractService.Iface
{
    @Override
    public FileDataRsp uploadFile(FileDataReq filedata) throws TException
    {
        System.out.println("uploadFile function has been called.");
 
        // 写到文件
        String filePath = "./2.txt";
        try
        {
            java.io.File file = new java.io.File(filePath);
            FileOutputStream fos = new FileOutputStream(file);
            FileChannel channel = fos.getChannel();
            channel.write(filedata.buff);
            channel.close();
        }
        catch (Exception x)
        {
            x.printStackTrace();
            return new FileDataRsp(1, "failed");
        }
        return new FileDataRsp(0, "successeed");
    }
}
