Folder dist with build files must be inside this server folder.

Generate build
GOOS=linux GOARCH=amd64 go build 

scp -i ~/.ssh/pv-alert-03-09-18.pem ./go-frontend-server ubuntu@ec2-18-205-247-206.compute-1.amazonaws.com:/home/ubuntu/market/

scp ./go-frontend-server devel@172.16.17.118:/home/devel/market/

```
 [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/ubuntu/market/go-frontend-server
  User=root
  Group=root
  Restart=always

  [Install]
  WantedBy=multi-user.target
  ```