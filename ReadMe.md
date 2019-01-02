# sync-ip
一个本地ip同步器，go编写

## 应用场景
- 非固定公网ip的测试服务器
- 游戏服务器
- 任何需要同步临时ip的机器

## 运行环境
linux/windows/osx

## 使用
服务端
```shell
git clone https://github.com/zzlhr/sync-ip.git
cd sync-ip/server
go build .
./server
```


客户端
```shell
git clone https://github.com/zzlhr/sync-ip.git
cd sync-ip/client
go build .
vi ./conf #修改server为自己服务器的ip
./client
```