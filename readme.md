### 本地开发环境
#### 1、 安装 yarn
```go
npm install -g yarn
yarn -v
```

#### 2、 安装 air
```go
go install github.com/air-verse/air@latest
air -v
```

### 启动服务

#### 1、启动前端服务

```go
// 切换到 cms 目录
cd cms
// 安装依赖
yarn install
// 启动服务
yarn dev
```

#### 2、启动后端服务
将 `.env.example` 复制为 `.env`， 从 `database/backups/` 目录手动导入 `sql` 文件到你的数据库中， 然后修改 `.env` 文件里端口号、数据库、Redis等配置信息

```go
// 切换到 service 目录
cd service
// 安装依赖
go mod tidy
// 启动服务
./serve.sh
```

