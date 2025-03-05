### 本地开发环境

####  安装 air
```go
go install github.com/air-verse/air@latest
air -v
```

#### 启动服务
将 `.env.example` 复制为 `.env`， 从 `database/backups/` 目录手动导入 `sql` 文件到你的数据库中， 然后修改 `.env` 文件里端口号、数据库、Redis等配置信息

```go
// 切换到 service 目录
cd service
// 安装依赖
go mod tidy
// 启动服务
./serve.sh
```

访问 http://127.0.0.1:3000/api 如果返回 `恭喜您 API 服务启动成功~ `说明服务启动成功

