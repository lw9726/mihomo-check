package config

const DefaultConfigTemplate = `
# 是否显示进度
print-progress: true

# 并发线程数
concurrent: 16

# 检查间隔(分钟)
check-interval: 30

# 超时时间(秒)
timeout: 5

# 下载测试大小(MB)
download-size: 20

# 上传测试大小(MB)
upload-size: 20

# 保存方法
save-method: r2

# 将测速结果推送到Worker的地址
worker-url: https://example.worker.dev

# Worker令牌
worker-token: 1234567890

# 订阅地址
sub-urls:
  - https://example.com/sub.txt
  - https://example.com/sub2.txt

# IP信息配置
ip-info:
  # IP查询API
  api-url:
    - http://ifconfig.me
    - http://ip.sb
    - http://ifconfig.es
    - http://ipinfo.io/ip
    - http://ipecho.net/ip
    - http://ident.me
    - http://eth0.me
    - http://ipaddr.site
    - http://ipaddress.sh
    - http://l2.io/ip
    - http://tnx.nl/ip
    - http://wgetip.com
    - http://ip.tyk.nu
    - http://curlmyip.net
    - http://ipcalf.com
    - http://checkip.amazonaws.com

  # IP数据库下载地址
  ipdb-url: https://cdn.jsdelivr.net/npm/openipdb.ipdb@2025.1.4/openipdb.ipdb
`