# Backend

The main purpose of this part is to serve information and plugins list, configured by the back-office.

## How to run the backend

### With Docker

You can run the backend server with the provided docker image :

```shell
docker container run \
  -d -p 8090:8090 \
  -v $(pwd)/data:/backend/data \
  -v $(pwd)/plugins:/backend/plugins \
  -v $(pwd)/medias:/backend/medias \
  -v $(pwd)/logs:/backend/logs \
  zenika/marcel-backend
```

 - `/backend/data/` is a volume that contains persitent data.  It should contains the plugin catalog in `plugins.json` (see [plugins README](../plugins)). The backend will generate a `medias.json` if it doesn/t exist.
 - `/backend/plugins/` is a volume that contains all the plugins files. Each plugin registered in the `plugins.json` catalog should have a folder with its `eltName` name.
 - `/bakcend/medias/` is a volume that should not be modified. It stores internal copies and data for each medias and plugins.
 - `/backend/logs/` is the volume containing the backend logs file.

### With compiled exectubale

Once built, you can run the exectuable as a background job :

```shell
./marcel-backend &
```

The backend expect a specific file architecture for the working directory :

```
working_directory
  |__ data/
  |   |__ plugins.json
  |   |__ medias.json
  |
  |__ medias/
  |   |__ ...
  |
  |__ plugins/
  |   |__ plugin1/
  |   |   |__ frontend/
  |   |       |__ index.html
  |   |       |__ style.html
  |   |
  |   |__ plugin2/
  |   |   |__ frontend/
  |   |   |   |__ index.html
  |   |   |__ backend/
  |   |       |__ docker_image.tar
  |   |
  |   |__ ...
  |
  |__ marcel.log
```

 - `data/` folder is used to store persistente data. It should contains the plugin catalog in `plugins.json` (see [plugins README](../plugins)). The backend will generate a `medias.json` if it doesn/t exist.
 - `medias/` folder should not be modified. It stores internal copies and data for each medias and plugins.
 - `plugins/` folder contains all the plugins files. Each plugin registered in the `plugins.json` catalog should have a folder with its `eltName` name.
 - `marcel.log` is the default location for logs.


## How to build the backend

### Dependencies

Install dependencies with [dep](https://github.com/golang/dep) :

```shell
go get -u github.com/golang/dep/cmd/dep
dep ensure -vendor-only
```

### Build server

You can build the server by running :

```shell
go build -o marcel-backend
```

You should then be able to run the server :

```shell
./marcel-backend
```

By default, the logs can be seen in the file `marcel.log` in the working directory. You can change this default :
```shell
export MARCEL_LOG_FILE="path_to_log_file" # defaults to $PWD/marcel.log
```

### Build cross architecture :

``` shell
env GOOS=linux GOARCH=arm go build -o marcel-backend-arm
```
List of all GOOS and GOARCH values : https://golang.org/doc/install/source#environment

### Build for Docker :

A Go executable needs some external library missing from alpine docker image. To compile a staticly linked executable usable for building the docker image, run :

``` shell
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo
```

## Credits

* [Gorilla Mux](https://github.com/gorilla/mux)
* [MapStructure](https://github.com/mitchellh/mapstructure) from Mitchellh
