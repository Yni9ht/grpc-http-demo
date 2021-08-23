# grpc-http-demo

## 前置条件
1. 安装 ETCD 集群，使用项目根目录中的 `etcd-docker-compose.yml` 快速启动一个 ETCD 集群

## Run
1. 在项目根目录执行下列命令，用于生成相关 pb 文件。
    ```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. proto/user.proto
    ```
   
2. 在 `server` 包中启动 main 函数，用于启动 server 端，该函数会启动一个 grpc 服务和 http 服务。

3. 在 `client` 包中启动用于测试调用 grpc 接口，server 端的控制台会打印相关内容。

4. 可通过 postman 访问服务端地址，请求路径：`127.0.0.1:8080/v1/auth/user/info`，server 端的控制台会打印相关内容。