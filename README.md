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
在浏览器中输入[http://\<ip>:\<port>/q/swagger-ui/](http://\<ip>:\<port>/q/swagger-ui/) （注意结尾有斜杠/），同时在Swagger ui界面的Explore栏目中输入[http://\<ip>:\<port>/q/service/\<package_name>.\<service_name>](http://<ip>:<port>/q/service/\<package_name>.\<service_name>)并点击Explore
![Alt text](/img/swagger.png)

## FAQ
如果访问/q/services出现报错`failed to decompress enc: bad gzipped descriptor: EOF`的报错说明部分依赖的proto文件生成的路径不对导致的，
比如:
- api/basedata/tag/v1/tag.proto
- api/basedata/article/v1/article.proto
当 **api/basedata/article/v1/article.proto** import "api/basedata/tag/v1/tag.proto"
时 ，生成的依赖为api/basedata/tag/v1/tag.proto 但是 tag.proto 生成的tag.pb.go文件中source是tag.proto（漏掉了api/basedata/tag/v1/）导致了依赖未找到

此时需要生成tag.pb.go时kratos proto client api/basedata/tag/v1/tag.proto 补全proto的路径，这样生成的tag.pb.go文件中source就是正确的(api/basedata/tag/v1/tag.proto)
