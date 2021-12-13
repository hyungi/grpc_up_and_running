# How to build
## 1. Build by shell command
```shell
# build server
$ cd /service
$ go build -v -o bin/server

# build client
$ cd /client
$ go build -v -o bin/client
```
## 2. Build by Makefile
Makefile is added you can easily generate proto file and build and remove artifact.

By this code, you can get server image at productinfo/{server|client}/bin
```shell
$ cd productinfo/{server|client}
$ make generate && make build
```

## Detail about build flags
### `-i`
In the book, they use `-i` flag. But it is deprecated.
```
The -i flag installs the packages that are dependencies of the target.
The -i flag is deprecated. Compiled packages are cached automatically.
```

### `-v`
Print the names of packages as they are compiled.

### `-o`
It means output, so you can define location and name of build result.

# How to run
## Run server
```shell
$ service/bin/server
```

## Run client
```shell
$ client/bin/client
```

## Server side message
```shell
2021/12/01 23:27:45 Product 248443e1-3363-4429-a443-bc37d06c78c0 : Apple iPhone 11 - Added.
2021/12/01 23:27:45 Product 248443e1-3363-4429-a443-bc37d06c78c0 : Apple iPhone 11 - Retrieved.
```

## Client side message
```shell
2021/12/01 23:27:45 Product ID: 248443e1-3363-4429-a443-bc37d06c78c0 added successfully
2021/12/01 23:27:45 Product: id:"248443e1-3363-4429-a443-bc37d06c78c0" name:"Apple iPhone 11" description:"Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode." price:699
```