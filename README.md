```azure
###增加GRPC调用接口
```
```
本地通过配置http代理 从世界各个地点 检测http服务的可用性
使用boltdb持久化数据 重新启动不需要重复配置数据
支持秒级别定时器
```

```
使用方式
 -d    是否删除 设置为是则删除所配置的url proxy
  -i int
        时间间隔 (default 10)
  -proxy string
        代理地址,号分隔eg:socks5://127.0.0.1:8000
  -url string
        地址多个地址逗号分隔eg:https://www.baidu.com
```
```
效果
$ ./httpMonitor.exe -i=2
执行监控: https://www.baidu.com
 -- success //为空为未配置代理输出
执行监控: https://www.baidu.com
 -- success
```