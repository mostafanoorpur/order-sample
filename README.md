# order-sample
**A service to handle order**

Stack:
- [Golang](https://golang.org/)

Libraries:
- [gorm](https://gorm.io/)  database orm
- [http]()
- [env package](https://github.com/spf13/viper)

step by step to run application :

* copy config.example.yaml and make a new file named config.yaml

* set database and redis user and password in (config.yaml) file

* in root folder
```
    make migrate
```
* start application

```
    make run
```

or you can run application by this command :

```
   docker-compose up --build
```