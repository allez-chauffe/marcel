# marcel

![Marcel](https://raw.githubusercontent.com/Zenika/marcel/master/marcel_banner.png)

marcel is a configurable plugin based dashboard system.

:construction: This README is still a work in progress...

## Create a plugin

Marcel is based on plugins, and we need you to complete the collection !

By convention, a plugin should have a name begin with `marcel-plugin-*` (`marcel-plugin-text` for example).
This ways, you can find a list of all available plugins by [searching them on github](https://github.com/search?utf8=%E2%9C%93&q=marcel%2Dplugin)

[See the marcel-plugin package to know more about plugin creation.](./node-packages/marcel-plugin)

## Build

`marcel` is composed of 3 parts:
  - `api` is the backend written in go.
  - `backoffice` is the single page app used to configure `marcel` and create medias.
  - `frontend` is the simgle page app actually displaying a media.

### Requirements

 - go > 1.11.0 (`marcel` is using go modules)
 - node > 9.0.0
 - npm > 5.0.0
  
### Bakcend
 
Buliding the backend is simple, you can just install the main go package :

```bash
$ go install
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
  - `/` : the backend (`backend/build`)
  - `/front` : the frontend (`frontend/build`)
  - `/api`: the backend (`localhost:8090` by default)
  
## Developpement
 
To have a working developpement environement, you have to run this 3 commands in seperated terminals :

```bash
$ cd backoffice && yarn && yarn start
$ cd frontend && yarn && yarn start
$ go build && ./marcel api --secure=false
```

You can then begin to modify sources. The backend is not compiled in watch mode, so you have to restart it manually. The backoffice and the frontend are liverloaded.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details
