# Backend

The purpose of this part is to give an serve a graphical administration interface. It allows to create, modify and delete medias with a graphical editor.

## How to run the back-office

### Form sources

Install dependencies :

```shell
  yarn
```

You can then create a production build :

```shell
yarn build
```

Once built, the back-office can be simply served by any http server :

```shel
serve -s ./build
```

The back-end and the front-end URI have to be configured in the `build/conf/config.json` file :

```json
{
  "backendURI" : "http://localhost:8090/api/v1/",
  "frontendURI" : "http://localhost/"
}
```

Note that the trailling slash of each URI is required.

### With docker

The back-office can be run with the provided docker image :

```shell
docker container run \
  -d -p 81:80 \
  -v $(pwd)/conf:/usr/share/nginx/html/conf
  marcel-back-office
```

The conf volume should contains the `config.json` file :

```json
{
  "backendURI" : "http://localhost:8090/api/v1/",
  "frontendURI" : "http://localhost/"
}
```

Note that the trailling slash of each URI is required.

## How to Contribute

### Run the developpement build

You can start the developpement server with live reloading :

```shell
yarn start
```

### Run the tests

You can run tests :

```shell
yarn test
```

or run the typechecker :
```shell
yarn flow
```

