# 安装logstash
```
mkdir -p /home/data/logstash/conf.d
scp -r logstash/logstash.yml /home/data/logstash/
scp -r logstash/suricata.conf /home/data/logstash/conf.d/
```


# 启动
```
docker run --name elasticsearch -d --restart=always -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" -e "network.host=localhost,127.0.0.1" --network=host elasticsearch:7.0.1

docker run -d --restart=always --net=host --name logstash -v /home/data/logstash/logstash.yml:/usr/share/logstash/config/logstash.yml -v /home/data/logstash/conf.d/:/usr/share/logstash/conf.d/ -v /var/log:/var/log logstash:7.0.1
```
