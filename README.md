# go-crud-demo
字节跳动青训营后端项目
# 极简抖音后端开发之路Day1

## 目标

1.  基本完成所有接口
2.  为后续的和他人合作完成项目做一个预演
3.  学习Go语言开发的一些规范
4.  学习一些小技巧
5.  整合目前自己所学的知识点
6.  丰富简历

## 使用工具确定

1.  工具语言Go语言
2.  Web框架采用Gin
3.  数据库目前采用MySQL，后续看情况是否使用Redis
4.  开发工具使用Goland
5.  数据库连接ORM工具采用gorm

## 开发环境搭建

1.  目录构建

   ├─Controller //控制层
   ├─Dao  //数据层
   ├─Middleware  //中间件层处理认证解析Token处理上文信息，为提供下文所需要的信息
   ├─public  //静态资源
   │  ├─mov  //视频资源
   │  └─pic  //封面图片资源
   ├─Service  //业务逻辑层
   │  ├─favorite
   │  ├─publish
   │  └─user
   └─Util  //自定义工具函数

2.  导包

```go
// go.mod
module go-crud-demo

go 1.20

require (
	github.com/aws/aws-sdk-go v1.44.317 // indirect
	github.com/bytedance/sonic v1.10.0-rc3 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.16.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/u2takey/ffmpeg-go v0.4.1 // indirect
	github.com/u2takey/go-utils v0.3.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/image v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/driver/sqlite v1.5.2 // indirect
	gorm.io/gorm v1.25.2 // indirect
)
```

## 极简抖音APP相关文档

1.  极简抖音App使用说明 - 第六届青训营版
   1.  提供Apk安装包
   2.  接口与场景说明
   3.  App使用说明
2.  抖音项目方案说明-第六届青训营后端项目
   1.  接口参数说明
3.  详细接口参数说明
