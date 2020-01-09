### 一次请求没问题，100万次可能就就有问题了

## 看redis连接数
1. netstat -nap | grep 6379 | wc -l 发现几万连接数

1. 重启ng /home/footstone/openresty/nginx/sbin/nginx -p /home/footstone/openresty/nginx -s reload
   再看连接数，正常。





