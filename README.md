A demo go app example，including integration with mongodb, redis, etc



## create test mongodb user
```
use go_app
db.createUser(
   {
     user: "go_app",
     pwd: "GoAppPass123",
     roles: [ "readWrite", "dbAdmin" ]
   }
)
```

## update dev constants

update `MongoDbHosts` in `constants/constants_dev.go'.

## build

exec `go build --tags dev`

## test

exec './go-app-example'