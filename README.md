# Learn Go By Example


## Go_Web 脚手架
- main.go：初始化，启动服务
- config：yaml、json等配置文件
- controllers：处理请求
- dao：负责数据与存储
- logger：日志文件
- logic: 业务逻辑
- model：模型
- pkg：第三方库
- router：路由

## Docker 启动MySQL、Redis

- 启动MySQL：`docker run -it --network=host mysql:8.0 mysql -h 0.0.0.0 camps_user -P 8086 -u root -p`
- 启动Redis：`docker exec -it 9268f9f58b0a redis-cli -h 0.0.0.0 -p 6379`