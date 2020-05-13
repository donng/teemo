teemo: 提莫虽小，五脏俱全

## 功能
- cron：每天定时抓取维基百科的“新闻动态”和“历史上的今天”
- api：提供抓取信息的接口

## 安装

1.复制配置文件，并修改其中的配置信息
```
cp conf/env.example.yaml conf/env.yaml
```
2.运行程序
```
go run main.go
```

## 依赖

- [gin](https://github.com/gin-gonic/gin)
- [goquery](https://github.com/PuerkitoBio/goquery)
- [gorm](https://github.com/jinzhu/gorm)
- [cron](https://github.com/robfig/cron)