# swagger-api
## Quick Start

在项目中引入openapiv2

```go
import	"github.com/go-kratos/swagger-api/openapiv2"

h := openapiv2.NewHandler()
//将/q/路由放在最前匹配
httpSrv.HandlePrefix("/q/", h)
```
启动应用后,在浏览器中输入 [http://\<ip>:\<port>/q/services](http://ip:port/q/services),找到需要渲染的\<package_name>.\<service_name>
![Alt text](/img/services.png)
在浏览器中输入[http://\<ip>:\<port>/q/swagger-ui/](http://\<ip>:\<port>/q/swagger-ui/)，同时在Swagger ui界面的Explore栏目中输入[http://\<ip>:\<port>/q/service/\<package_name>.\<service_name>](http://<ip>:<port>/q/service/\<package_name>.\<service_name>)并点击Explore
![Alt text](/img/swagger.png)
