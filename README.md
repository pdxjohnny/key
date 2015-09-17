Key
---

This package strives to be an easy to use wraper on golangs crypto

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
go build -o key_linux-amd64 -tags netgo *.go
# Or
./script/build
# My favorite
SKIP_BUILD=1 SKIP_IMAGE=1 ./script/build -osarch="linux/amd64"
```
> `-tags netgo` will help you achieve static binaries :)

Running
---

```bash
john@localhost:key$ ./key_linux-amd64 gen
john@localhost:key$ ./key_linux-amd64 encrypt --message="Hello World" > /tmp/t
john@localhost:key$ ./key_linux-amd64 decrypt --message='/tmp/t'
Hello World
john@localhost:key$

```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
