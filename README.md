
# 项目结构
- **certs**：用于存放 SSL/TLS 证书和密钥，用于 HTTPS 服务的安全加密。 
- **clients**：包含与外部服务或数据库的客户端代码，用于连接和交互。
- **controllers**：包含控制器的实现，负责处理 HTTP 请求，执行业务逻辑，并返回响应（MVC架构）。
- **dao（Data Access Object）**：用于数据访问对象，包含与数据库交互的代码，如执行 CRUD（创建、读取、更新、删除）操作。
- **global**：包含全局配置或全局变量，这些配置或变量在应用程序的多个部分中使用。
- **log**：用于存放日志相关的代码或配置，如日志记录器的初始化和日志格式的定义。
- **middleware**：处理 HTTP 请求和响应的函数，如身份验证、日志记录、请求限制等。
- **models**：包含定义数据结构和数据库模型的代码，与数据库表结构对应。
- **resp**：用于存放响应相关的代码，如标准化的响应格式或错误处理。
- **routers**：包含定义应用程序路由的代码，即 URL 路径到处理函数的映射。
- **utils**：存放工具函数或实用程序代码，这些代码可以在应用程序的多个部分中复用。
- **websites**：这个文件夹的用途不太明确，可能用于存放静态网站文件，如 HTML、CSS、JavaScript 文件，或者用于存放与网站相关的配置。

---
# 待办
- [ ] 路由设置
  - api 接口调整，删减
- [ ] 数据模型设置
- [ ] 日志记录
  - 日志输出设置有误，需要调整
- [ ] 中间件调整
- [ ] 数据库根据模型进行迁移


---
# 功能
- 用户：
  - 登录
  - 退出
- 商品
  - 查看
  - 选中商品
  - 加入购物车