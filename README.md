# marcel

![Marcel](https://raw.githubusercontent.com/Zenika/marcel/master/marcel_banner.png)

Marcel is a configurable plugin based dashboard system.

:construction: This README is still a work in progress...

## Create a plugin

Marcel is based on plugins, and we need you to complete the collection!

By convention, a plugin should have a name begin with `marcel-plugin-*` (`marcel-plugin-text` for example).
This way, you can find a list of all available plugins by [searching them on gitHub](https://github.com/search?utf8=%E2%9C%93&q=marcel%2Dplugin)

[See the marcel-plugin package to know more about plugin creation.](./node-packages/marcel-plugin)

## Build

`marcel` is composed of 3 parts:
  - `api` is the backend written in go.
  - `backoffice` is the single page app used to configure `marcel` and create medias.
  - `frontend` is the single page app actually displaying a media.

### Requirements

 - go > 1.11.0 (`marcel` is using go modules)
 - node > 9.0.0
 - npm > 5.0.0
  
### Backend
 
Building the backend is simple, you can just install the main go package :

```bash
$ go install ./cmd/marcel
```

This will make the `marcel` command available (if your go bin folder is in your PATH)

### Backoffice and Frontend

The backoffice and the frontend are both regular React application. To build them, go to their respective folder and run

```bash
$ npm i && npm run build
# or with yarn
$ yarn && yarn build
```

## Usage

The backend can be launched with the `marcel` command line :

```bash
$ marcel api
```

This will serve the api on the default port `8090`.

You should then serve `api`, `backoffice` and `frontend` behind reverse proxy and serve this routes :
  - `/` : the backoffice (`pkg/backoffice/build`)
  - `/front` : the frontend (`pkg/frontend/build`)
  - `/api`: the backend (`pkg/api`)

`localhost:8090` by default
  
## Development
 
To have a working development environment, you have to run this 3 commands in separated terminals :

```bash
$ go build ./cmd/marcel && ./marcel api --secure=false
$ cd pkg/backoffice && yarn && yarn start
$ cd pkg/frontend && yarn && yarn start
```

You can then begin to modify sources. The backend is not compiled in watch mode, so you have to restart it manually. The backoffice and the frontend are live-reloaded.

Another solution is to use the `standalone` mode if you want a quick launch :

```bash
$ go build ./cmd/marcel && ./marcel standalone
```

(don't forget to save admin password displayed in logs :-))

or the `demo` mode if you just want to play with it :

```bash
$ go build ./cmd/marcel && ./marcel
```

If you want to explore the bolt database, you can use the tool [boltdbweb](https://github.com/evnix/boltdbweb)

```bash
$ go install go get github.com/evnix/boltdbweb@latest
$ boltdbweb --db-name=marcel.db --port=<port>[optional] --static-path=<static-path>[optional]
```


## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details
