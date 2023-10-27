# Frame

```shell
# 生成ent资源文件
go generate ./ent
# 生成swagger文档
go run -mod=mod github.com/swaggo/swag/cmd/swag@latest init -g router/new.go
```

## Entity SQL生成

### Step 1

`go run -mod=mod entgo.io/ent/cmd/ent new XXX(表名, 大驼峰)` 生成新表模板

修改 `ent/schema` 目录下的生成文件

`go generate ./ent` 生成 `entity` 增删改查代码.

业务代码完成需要运行测试时进行下一步.

### Step 2

```shell
atlas migrate diff xxxxxx(生成对应名称的sql文件) \                                        
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "mysql://root:root@localhost:3306/dev"
```

### Step 3

检查是否需要修改SQL, 修改sql完成后执行

```shell
# 重新生成hash
atlas migrate hash --dir "file://ent/migrate/migrations"
```

### Step 4

```shell
# 检查迁移
atlas migrate lint \
  --dev-url "mysql://root:root@localhost:3306/dev" \
  --dir "file://ent/migrate/migrations" \
  --latest 1
# 应用迁移文件
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://root:root@localhost:3306/test"
```
