`go run main.go http://yahoo.com http://google.com http://romtin.github.io http://undefined_url.com`

```
4.20s    4993 http://romtin.github.io
Get http://undefined_url.com: dial tcp: lookup undefined_url.com on 192.168.11.1:53: no such host
5.69s   11997 http://google.com
7.83s  511658 http://yahoo.com
7.83s elapsed
```