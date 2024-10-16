# 安装服务
```
scp -r system/*.service /usr/lib/systemd/system/
systemctl daemon-reload

systemctl enable my-zero-api.service
systemctl enable my-zero-loginsvc.service
systemctl enable my-zero-usersvc.service


systemctl start my-zero-loginsvc.service
systemctl start my-zero-usersvc.service
systemctl start my-zero-api.service
```