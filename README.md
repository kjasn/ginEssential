# ginEssential

## 用到的第三方工具：

```bash
go get -u github.com/gin-gonic/gin # gin框架
go get -u gorm.io/gorm  # 通过gorm来操作数据库
go get -u gorm.io/driver/mysql v1.5.1 # 数据库驱动,用于连接数据库
go get github.com/spf13/viper   # config组件 -- viper 用于读取配置文件
```



## 遇到的问题：

1. 我在这个独立的项目设置了`main`分支保护，我在其他分支提交PR并向`main`分支合并时需要审批。但是，我作为这个PR的发起人而无法给自己审批...
到处找解决方法都没有找到，好像可以强制合并，我没有试过，因为如果每次都强制合并那开启分支保护不就没意义了？实在找不到别的办法，看到有个此次PR不需要审批的选项就勾上了（跟强制合并差不多）。
