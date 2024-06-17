# 发布 Go 包到 GitHub

## 发布

仓库初始化成

```
git mod init github.com/code-elastic/open-api-client-sdk-go
```

提交代码

```
git add 
git commit
```

**打标签**

```
git tag -a v1.0.0 -m "Initial release"
git push origin v1.0.0
```



在 GitHub 仓库的 release 中选择一个 tag 进行发布，写相关发布信息，并 “ **Set as the latest release** ”





## 使用/更新

```
go get -u github.com/code-elastic/open-api-client-sdk-go
```

go mod tidy