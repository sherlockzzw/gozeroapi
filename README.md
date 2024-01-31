# apiClipFilm 一剪成片相关功能api接口


## 全局说明：

1. 这是一剪成片相关功能api接口项目


## 开发说明

1. 项目把根目录下的`.api`文件移动到`/api`目录下，并更名为`run.api`
2. 新增模块需要在`/api`目录下新增一个`.api`文件，并在`run.api`文件中注入`import "user.api"`
3. 新增接口根据官方文档或者现有的注释使用，注意⚠️：不是自己的不要动，执行后把handler和logic中的数据覆盖
4. 执行命令
```shell
cd api
# 构造所有的服务
goctl api go -api ./run.api -dir ../
# 构造某一个服务
goctl api go -api ./user.api -dir ../
```


## 响应规范

1. 不能把报错信息响应到前端
2. 报错需要使用错误码及错误信息响应
3. 可以自己构造也可以使用现有的已构造好的
4. 现有的是在main方法中执行`internal/global/errorMsg`包中的`errorMsg.New()`方法
5. 需要在已有的`internal/global/errorMsg`包中自定义现有项目所需的错误码变凉，并使用add方法注册到错误信息服务中
6. 在响应中使用：
```shell
# 成功
result.ResultSuccess(r.Context(), w, 123)
# 错误
result.ResultFail(r.Context(), w, result.Error, err)
```

## 文档

官方文档：https://go-zero.dev/docs/tasks/http/hello-world


