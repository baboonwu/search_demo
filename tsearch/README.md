# tapi_go
The API service of TOWN. Go Version.

# 1 Install
确保安装了 git
## 1.1 Install golang 1.4.2 (64bit)
### Mac下安装：
在官网http://golang.org/dl 直接下载安装包安装即可。下载pkg格式的最新安装包，直接双击运行，一路按照提示操作即可完成安装。

### 配置
设置你的本地代码目录，比如 /Users/ubuntu/code/go

export GOPATH=/Users/ubuntu/code/go

## 1.2 Get Code

### 1.2.1 tapi_go
代码地址：https://github.com/augmn/tapi_go

建立你的代码目录
/Users/ubuntu/code/go/src

### 1.2.2 三方库
- go get github.com/ant0ine/go-json-rest/rest
- cd github.com/AdRoll & git clone https://github.com/augmn/goamz
- go get github.com/garyburd/redigo/redis
- go get github.com/satori/go.uuid


### 1.2.3 Install Redis: [Download](http://redis.io/download)

	```bash
	$ wget http://download.redis.io/releases/redis-2.8.19.tar.gz
	$ tar xzf redis-2.8.19.tar.gz
	$ cd redis-2.8.19
	$ make

	$ src/redis-server
    ```

## 1.3 Build
go build main.go

## 1.4 Run

go run ./main.go -cfg config/local.cfg

# 2 IDE
## 2.1 Download [IDEA Free Version](https://www.jetbrains.com/idea/download/)
## 2.2 Install [plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/Documentation#installing-the-plugin)
## 2.3 Debug uses [delve](https://github.com/derekparker/delve)


