# Go time manager

## 主要功能
- gorm对接mysql，自动迁移创建表，gen自动生成curd
- 使用gin web框架，封装响应，中间件等
- 支持https和http2
- redis接入
- 发送邮箱验证码，配合redis管理验证码
- jwt生成token，配合redis刷新token过期时间
- 自动生成swagger接口文档（knife版本）
- 高性能日志工具zap，配合lumberjack自动分割日志（按大小和按天分割）
- 利用flag进行多环境配置
- yaml配置文件管理
- 利用makefile进行项目打包
