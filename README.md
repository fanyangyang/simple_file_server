# simple_file_server
用 GO 实现的文件服务器，局域网内暴露服务器文件供外部用户快速下载

能够指定启动端口

能够指定路径

```bash
go run main.go --port 1235 --path="/Users/fanyangyang"
```


## TODO 

- [ ] 优雅退出
- [ ] 输出帮助文档