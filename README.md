
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
- [ ] User auth
- [ ] Login
- [ ] Register
- [ ] Storage
- [ ] Benchmark