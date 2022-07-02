### 『仿知乎专栏』项目是高仿知乎专栏的后端项目，基于 Gin 框架从零到一到上线

### 所有路由

| 请求方法 | API 地址                                | 说明                        |
| -------- | --------------------------------------- | --------------------------- |
| POST     | /api/v1/auth/login/using-phone          | 短信 + 手机号登录           |
| POST     | /api/v1/auth/login/using-password       | 手机号、用户名、邮箱 + 密码 |
| POST     | /api/v1/auth/login/refresh-token        | 刷下 Token                  |
| POST     | /api/v1/auth/password-reset/using-email | 邮件密码重置                |
| POST     | /api/v1/auth/password-reset/using-phone | 短信验证码密码重置          |
| POST     | /api/v1/auth/signup/using-phone         | 使用手机号注册              |
| POST     | /api/v1/auth/signup/using-email         | 使用邮箱注册                |
| POST     | /api/v1/auth/signup/phone/exist         | 手机号是否已注册            |
| POST     | /api/v1/auth/signup/email/exist         | email 是否已支持            |
| GET      | /api/v1/user                            | 获取当前用户                |
| PUT      | /api/v1/users                           | 修改个人资料                |
| PUT      | /api/v1/users/email                     | 修改邮箱                    |
| PUT      | /api/v1/users/phone                     | 修改手机号                  |
| PUT      | /api/v1/users/password                  | 修改密码                    |
| GET      | /api/v1/columns                         | 分类列表                    |
| POST     | /api/v1/columns                         | 创建分类                    |
| PUT      | /api/v1/columns/                        | 更新专栏                    |
| GET      | /api/v1/columns/:id/posts               | 文章列表                    |
| POST     | /api/v1/topics                          | 创建文章                    |
| PUT      | /api/v1/posts/:id                       | 更新文章                    |
| DELETE   | /api/v1/posts/:id                       | 删除文章                    |
| GET      | /api/v1/posts/:id                       | 获取文章                    |
| POST     | /api/v1/upload                          | 上传文件                    |
| GET      | /api/v1/links                           | 友情链接列表                |

### 第三方依赖

使用到的开源库：

- [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
- [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [cobra](https://github.com/spf13/cobra) —— 命令行结构
- [viper](https://github.com/spf13/viper) —— 配置信息
- [cast](https://github.com/spf13/cast) —— 类型转换
- [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
- [snowflake](https://github.com/mojocn/base64Captcha) —— 雪花算法
- [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [limiter](https://github.com/ulule/limiter/v3) —— 限流器
- [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
- [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
- [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理
- [faker](https://learnku.com/courses/go-api/1.17/finish-up/github.com/bxcodec/faker) —— 假数据填充
- [imaging](https://learnku.com/courses/go-api/1.17/finish-up/github.com/disintegration/imaging) —— 图片裁切

### 自定义的包

现在来看下我们自建的库：

- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- snowflake —— 雪花算法
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- oss —— 对象存储
- str —— 字符串处理
- verifycode —— 数字验证码

### 代码行数

GDColumn 项目总共有 6300 行代码（工具 [gocloc](https://github.com/hhatto/gocloc)）：

```go
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             115           1015            562           4963
JSON                             1              0              0            672
YAML                             2              0              1            467
XML                              4              0              0            129
BASH                             1              8              0             90
TOML                             1              5             23             28
SQL                              1              0              1              1
-------------------------------------------------------------------------------
TOTAL                          125           1028            587           6350
-------------------------------------------------------------------------------
```



\- ***\*前端演示站点：[**http://bitepig.aa12.cool/*]**(http://bitepig.aa12.cool/)***\***

\- 后端 API 查询站点：[http://bitepig.aa12.cool/swagger/index.html](http://bitepig.aa12.cool/swagger/index.html)

**## 演示信息**

账号密码: test@test.com / 123456
