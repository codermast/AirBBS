# Air Blog

## 项目介绍

本项目是基于 Golang 语言开发的一款前后端分离的论坛社区管理系统，作为本人 Golang 学习中的一个完整项目。具体的技术栈如下：

- Golang 1.22.2
- Gorm
- MySQL
- Vue3
- Naive UI

## 项目部署

### Docker部署

1. 下载 Docker 镜像

```shell
docker pull codermast/airbbs:latest
```

2. 启动 Docker 容器

```shell
docker run -p 8080:8080 codermast/airbbs:latest
```
3. 访问项目

```http request
http://localhost:5174
```

### 编译部署

本仓库提供 Windows、MacOS、Linux 下的 Arm64 和 X86_64 的后端编译版本。

### 源码部署

1. 克隆仓库到本地

```shell
git clone https://github.com/codermast/AirBBS
```

2. 修改数据库信息

```yaml
#  数据库配置信息
database:
  host: localhost
  port: 3306
  dbname: airbbs
  username: codermast
  password: 123456
```

3. 启动后端

```shell
go run main.go
```

4. 启动前端

```shell
# 先安装依赖
npm install 

# 启动项目
npm run dev
```

5. 访问项目

```http request
http://localhost:5174
```