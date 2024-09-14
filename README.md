

### doc
https://entgo.io/docs/getting-started/

### 步骤
```bash
1. go mod init entdemo
2. go run -mod=mod entgo.io/ent/cmd/ent new User
3. vim bala bala: add field define and "entgo.io/ent/schema/field" in user.go
4. go generate ./ent
5. add start.go
6. run
```

### 执行sql语句
```bash 
0. 需要引入sql/execquery特性
1. 修改[generate.go], 添加--feature sql/execquery
2. go generate ./ent
3. run
```
