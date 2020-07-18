learning 学习目录

conf：用于存储配置文件   
middleware：应用中间件    
models：应用数据库模型  
pkg：第三方包    
routers 路由逻辑处理  
runtime：应用运行时数据

###阅读go.mod文件
#####新增 replace 配置项 因为开始以下模块没推送到远程，无法下载下来，因此需要使用replace，将其指定读取本地的模块路径，这样就可以解决本地模块读取问题
replace (
    github.com/nearbyren/gomodule//pkg/setting => ~/go-application/gomodule/pkg/setting
    github.com/nearbyren/gomodule//conf    	  => ~/go-application/gomodule/pkg/conf
    github.com/nearbyren/gomodule//middleware  => ~/go-application/gomodule/middleware
    github.com/nearbyren/gomodule//models 	  => ~/go-application/gomodule/models
    github.com/nearbyren/gomodule//routers 	  => ~/go-application/gomodule/routers
)


v1/tag.go           类似java Control层
models/tag.go       类似java Dao层
setting.go          类似java 读取配置启动信息
models.go           类似java 链接数据库操作
router.go           类似java 的注解服务注入

PUT 访问 http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1 ，查看 code 是否返回 200
DELETE 访问 http://127.0.0.1:8000/api/v1/tags/1 ，查看 code 是否返回 200

在本机执行curl 127.0.0.1:8000/api/v1/tags，
正确的返回值为{"code":200,"data":{"lists":[],"total":0},"msg":"ok"}



POST：http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
GET：http://127.0.0.1:8000/api/v1/articles
GET：http://127.0.0.1:8000/api/v1/articles/1
PUT：http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
DELETE：http://127.0.0.1:8000/api/v1/articles/1

