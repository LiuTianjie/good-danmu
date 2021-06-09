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