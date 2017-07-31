# Backend

The main purpose of this part is to serve information and plugins list, configured by the back-office.

## How to build the project

```go
go get github.com/rs/cors
go get -u github.com/gorilla/mux
go get github.com/mitchellh/mapstructure
```

```shell
docker pull quay.io/goswagger/swagger
alias swagger="docker run --rm -it -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"

export MARCEL_LOG_FILE="path_to_log_file" # defaults to $PWD/marcel.log
```

Build cross architecture :

``` shell
env GOOS=linux GOARCH=arm go build -o ./bin/MARCEL
```
List of all GOOS and GOARCH values : https://golang.org/doc/install/source#environment

Once every is done, you can launch the server with :

```shell
go run main.go
```

###Swagger.json generation

Install library :
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

Generate the API doc :
```
swagger generate spec -o ./swagger.json
```

Run Swagger UI on port 3000

### How to develop a plugin
All plugins have 2 main folders and 1 description file :
```
 plugin_name
  |__front\
  |   |__ index.html
  |   |__ ...
  |
  |__back\
  |   |__ docker_image.tar
  |
  |__description.json
```

 * Each attribute for the backend will be passed to the docker image as a environment variable. 
 * For each instance of a launched plugin, a Docker Volume is created at ```/tmp``` on the container
 
 To get a full example, please have a look at the "Google Agenda" plugin here: 
 https://github.com/Zenika/marcel-plugin-calendar

## Credits

* [Gorilla Mux](https://github.com/gorilla/mux)
* [MapStructure](https://github.com/mitchellh/mapstructure) from Mitchellh
* [Go-Swagger.io](https://goswagger.io)