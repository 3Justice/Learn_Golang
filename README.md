# Learn_Golang
learning process of Golang

### Go工程初始化环境
#### 步骤1：设置项目结构
```sh
mkdir myproject
cd myproject  
```
  
#### 步骤 2：初始化Go模块
```sh
go mod init example.com/myproject
```  
例如
```sh 
go mod github.com/yourusername/myproject
```

#### 步骤 3：编写代码
根据需求编写合适的代码

#### 步骤 4：下载和管理依赖包
在代码中导入任何需要的第三方包
```sh
go get github.com/some/package
```  
自动清理和更新并下载所有依赖
```sh
go mod tidy
```  
#### 步骤 5：运行代码
##### 运行代码：
```sh
go run main.go
```  

##### 编译并生成可执行文件（可选）：
```sh
go build -o myproject
./myproject
```  