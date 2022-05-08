package main

/*


如果key是短字符串，value也是字符串，100 字节。
如果key是字符串， value 是包含5个字段的值，200字节。

导出查看服务器的所有keys :
echo  "KEYS *\n" | redis-cli -h 192.168.80.206 -p 6379 -a 111111 > keys.redis
计算每个key占用的内存：
for key in `cat keys.redis` ; do info=`echo "debug object $key" | redis-cli -h 192.168.80.206 -p 6379 -a 111111 | grep -o "serializedlength\:[0-9]*" | cut -d ":" -f2 | tr -d '\n'`; echo $info": "$key  >> keysize.redis; done;
对生成结果进行排序，查看内存使用大的key（这边列出前面100个）：
cat keysize.redis | sort -nr -k 1 | head -n 100

 */




