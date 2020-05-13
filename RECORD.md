# 项目创建过程

## 1. 项目初始化

```bash
# 1. go module 初始化项目
go mod init github.com/hotttao/goalgo

# 2. 下载 cobra
go get -u github.com/spf13/cobra/cobra

# 3. 初始化项目
cd goalgo
cobra init --pkg-name github.com/hotttao/goalgo

# 4. 添加子命令
cobra add sort
cobra add quick -p 'sortCmd'
cobra add binary

# 5. 查看帮助
go run main.go -h
go run main.go sort -h
```
