## go-gin-mgo-demo

A demo CRUD application in golang using the popular gin-gonic framework

## Development

1. Clone the (forked) repo.
2. Then, run

  ```sh
  $ go get github.com/codegangsta/gin
  ```

  [codegangsta/gin](http://github.com/codegangsta/gin) is used to to automatically compile files while you are developing

3. Run

  ```sh
  $ go get && go install && PORT=7000 DEBUG=* gin -p 9000 -a 7000 -i run # or make dev
  ```

  Then visit `localhost:7000`

- MONGODB_URL can be configured by setting the env variable.
- PORT can also be configured by setting the env variable

```sh
$ export MONGODB_URL=mongodb://
$ export PORT=7000
```

[godep](https://github.com/tools/godep) is used for dependency management. So if you add or remove deps, make sure you run `godep save` before pushing code. Refer to its documentation for more info on how to use it.

## Usage

```sh
$ go get github.com/madhums/go-gin-mgo-demo
$ PORT=7000 go-gin-mgo-demo # should start listening on port 7000
```

#### Credits

Thanks to all the dependent packages
