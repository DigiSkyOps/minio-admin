

# minio-admin

## 项目介绍
minio-admin 是一个用于 minio 用户及策略的管理工具，由 minio 原生的命令行方式变为操作简单方便的UI界面方式,该项目为前后端分离项目，前端vue，后端golang gin。
  
## 已有功能介绍
- 用户添加、删除、禁用、绑定策略、批量删除用户、批量禁用用户。
- 策略添加、删除、修改、批量删除，添加策略时支持自定义编写 json 格式的策略内容和按照 UI 界面点击生成策略内容。


## 项目结构
- 前端源码是该目录下的 minio-admin-fed，前端详细介绍地址：https://github.com/DigiSkyOps/minio-admin/minio-admin-fed
- 后端源码是该目录下的 minio-admin-platform，后端详细介绍地址：https://github.com/minio-admin-platform


## 项目部署
本项目使用容器部署，可以直接使用docker run 方式或者跑在 kubernetes  上，下面简单介绍 docker run 的部署方式。
- 构建部署 minio-admin-platform
```
$ cd minio-admin-platform

# 替换如下命令中 minio 的 ACCESSKEY、SECRETKEY、MINIO_ENDPOINT 变量
$ docker build -t minio-admin-platform --build-arg ACCESSKEY="xxx" --build-arg SECRETKEY="xxx"  --build-arg MINIO_ENDPOINT="xxx"  .

# 此处映射的本地端口为 9000 
$ docker run -itd --name minio-admin-platform -p  9000:9000 minio-admin-platform 
```
- 构建部署  minio-admin-fed
```
$ cd minio-admin-fed
# 将如下命令中的 PLATFORM_URL 变量值替换为上一步中已经部署的  minio-admin-platform 的地址
$ docker build -t minio-admin-fed  --build-arg PLATFORM_URL="xxx" .

$ docker run -itd --name minio-admin-fed -p  80:8000   minio-admin-fed 
```
- 测试浏览
```

使用浏览器打开前端应用 minio-admin-fed 的地址即可。

```

## 其他参考文档

-  minio 官网：https://min.io/
- 开源地址：https://github.com/minio/
- 中文文档：https://docs.min.io/cn/
