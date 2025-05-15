# GoUtils

#### 介绍
集合了一些 Go api开发中常用的工具和方法。

#### 软件架构
src
|--utils  工具类
|--auth  鉴权类
|--datetime  日期时间类
|--types  公共类（结构体）


#### 安装教程


#### 使用说明

1.  auth工具类：
    * 生成jwt
    * 解析jwt
    * md5加密
    * sha256加密
2.  datetime工具类：
    * 获取当前时间的零点时间，返回time.Time类型
    * 计算某一时间零点与当天的零点时间差，返回小时数
    * 计算某一时间到下一零点的时间差，返回time.Duration类型
3.  utils工具类：
    * 检测是否为合格手机号
    * 生成固定长度的随机字符串或数字

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

