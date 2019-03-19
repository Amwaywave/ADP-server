# ADP (API Document Platform) Server

API Document Platform 的服务端。


## 项目结构

```text
├── cmd                 # cli
│   └── adp
├── conf                # configs
├── models              # models
└── pkg
    └── languages       # 语言相互转换的实现
    
6 directories
```

## 功能

*还需细化*

- [ ] 接口相关
  - [x] 接口结构
  - [ ] 多语言相互转换
    - [x] Yaml
    - [ ] Java
    - [ ] TypeScript
  - [ ] 支持Fork
  - [ ] 支持Merge
  - [ ] 版本控制
  - [ ] 单点管理锁
  - [ ] 支持Mock
  - [ ] 反向创建（通过数据生成接口）
  - [ ] 导入、导出

- [ ] 用户相关
  - [ ] 用户分组（组织）
  - [ ] 用户登录
  - [ ] 用户注册
  - [ ] 第三方用户认证
    - [ ] Github
    - [ ] Gitlab
  - [ ] 用户权限

- [ ] 其他
  - [ ] 提供Hook

## 接口结构

```yaml
project: adp (API Document Platform)
author: 
  userName: 用户名
  email: 邮箱
name: 用户登录接口
description: 用户登录接口（系统用户或第三方登录）
type: POST
endpoint: /sessions
request:
  headers:
    contentType: application/json
  params:
    userName: 
      description: 用户名
      type: string
      checks: 
        - rule: len < 5
          pass: 用户名符合要求
          reject: 用户名长度不符合要求
    password:
      type: string
      checks:
        - rule: len < 6
          pass: 密码符合要求
          reject: 密码长度不符合要求
response:
  status: 200
  type: json
  body:
    message: 登录成功
    code: 0
    data:
      loginToken: 
        type: string
        description: 用户登录标识
```
