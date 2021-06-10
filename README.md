<!--
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:56:06
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 21:53:08
-->

### good-danmu
good-danmu is a high performance danmu server using gin framework, as a starter of Golang, it may help to better
understand how and why Golang is so suitable for high performance, concurrent programs.

### Category Guide
> /src Source codes.

> /src/api Apis to operate data, follow restful specifications.

> /src/handler Handlers.

> /src/init Initialize the program, db, router, etc.

> /src/middleware Middlewares, auth, interpretor, etc.

> /src/router Routers.

> /src/service Services, at present, it only contails danmu service, as the project expanded, we will add more service such as jwt, encryption services and etc.

> /src/utils Utils.

### 5 steps to build a highly available Redis Cluster.
Assuming you have installed docker and docker-compose.
#### 1. Pull the latest Redis docker image
```shell
sudo docker pull redis:6.0.6
```
#### 2. Creat 6 dictionaries to run 6 docker.
```shell
mkdir -p ./{6378..6383}/data
mkdir -p ./{6378..6383}/redis
```
#### 3. Configure the redis.conf
Here I just list the key options, after configuration, copy it
into different folds.
```editorconfig
# Open Cluster mode.
cluster-enabled yes
# The port should be the same as the current dictionary
# Here suppose we're in 6378 folder.
port 6378
# This option allow the slave node replace the master 
# immediately if the master is unreachable.
cluster-slave-validity-factor 0
# We don't need the redis run frontend.
daemonize no
# Enable the AOF.
appendonly yes
# Point out the Cluster timeout.
cluster-node-timeout 5000
# This configuration is option, if you want to access
# your Cluster remotely, write it and comment out the next line.
bind 0.0.0.0
# bind 127.0.0.1
```
#### 4. Run the docker-compose-redis.yaml scripts.
```shell
sudo docker-compose -f docker-compose-redis.yaml up -d
```
#### 5. Open Cluster mode
```shell
Here the default password I used is 1123581321
sudo docker run --rm -it goodsmileduck/redis-cli redis-cli -a `your_password` --cluster-replicas 1 --cluster create 10.112.187.89:6378 10.112.187.89:6379 10.112.187.89:6380 10.112.187.89:6381 10.112.187.89:6382 10.112.187.89:6383
```

### Todo list
- [x] Basic danmu structure
- [x] User auth
  - [x] Websocket auth
  - [x] Login
  - [x] Register
  - [ ] single place login
  - [ ] More complex auth with casbin
- [x] Storage
  - [x] user model
  - [x] danmu model
- [x] Use Redis to store danmu, cache login token.
  - [x] use redis to cache the login status
  - [x] storage
  - [x] redis storage AOF/RDB
  - [x] when a new client is online, search the exact channel's data, if the data is not in redis, do the search in mysql, and store it in redis, with time of 10 minutes.
- [ ] Benchmark