# swagger-api
## Quick Start

在项目中引入openapiv2

```go
import	"github.com/go-kratos/swagger-api/openapiv2"

h := openapiv2.NewHandler()
//将/q/路由放在最前匹配
httpSrv.HandlePrefix("/q/", h)
```

支持generator进行自定义配置
```go
h := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(true)))
```
更多配置参见 https://github.com/go-kratos/grpc-gateway/blob/master/protoc-gen-openapiv2/generator/option.go


启动应用后,在浏览器中输入 [http://\<ip>:\<port>/q/services](http://ip:port/q/services)，在顶栏右侧选框选取希望查看的服务名，即可浏览接口文档。
![select service](/img/swagger.png)

## FAQ
#### 1. 如果启动时顶栏选框未显示可选的服务名，或访问/q/services出现报错，`failed to decompress enc: bad gzipped descriptor: EOF`的报错说明部分依赖的proto文件生成的路径不对导致的，
比如:
- api/basedata/tag/v1/tag.proto
- api/basedata/article/v1/article.proto

当 **api/basedata/article/v1/article.proto** import "api/basedata/tag/v1/tag.proto"时 ，生成的依赖为api/basedata/tag/v1/tag.proto 
但是 tag.proto 生成的tag.pb.go文件中source是tag.proto（漏掉了api/basedata/tag/v1/）导致了依赖未找到

此时需要生成tag.pb.go时kratos proto client api/basedata/tag/v1/tag.proto 补全proto的路径，这样生成的tag.pb.go文件中source就是正确的(api/basedata/tag/v1/tag.proto)

#### 2. 给swagger api添加说明、examples
请参考[a_bit_of_everything](https://github.com/longXboy/grpc-gateway/blob/master/examples/internal/proto/examplepb/a_bit_of_everything.proto)，同时注意需要将[annotations.proto](https://github.com/go-kratos/kratos/blob/main/third_party/protoc-gen-openapiv2/options/annotations.proto)和[openapiv2.proto](https://github.com/go-kratos/kratos/blob/main/third_party/protoc-gen-openapiv2/options/openapiv2.proto)这两个proto文件复制到third_party/protoc-gen-openapiv2/options目录下
