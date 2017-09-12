# Frontend

This part of the project is what will be on the mirror. It's a lightweight html application with the inclusion of Polymer. At startup, it will load the list of plugins which have to be displayed, than load every plugin into the page and finally display them.

## Run

You can run the application with the provided image docker :

```shell
docker container run \
  -d -p 80:80 \
  -v $(pwd)/conf:/usr/share/nginx/html/conf \
  marcel-front-end
```

Note that you also have to configure the back-end and plugins URI in the `config.json` file in the `conf` volume :

```json
{
  "backendURL": "http://localhost:8090/api/v1",
  "pluginURL": "http://localhost:8090/api/v1"
}
```

Note that for now, the tow URL have to be the same.

## Installation

First, you have to fetch the bower components:

```bash
bower install
```

Then you have to copy the `conf/config.example.json` file into `conf/config.json` and change the addesses to specify the URL of the back-end and the plugins files.

Finally you just have to serve the static files with your favorite server like nginx or with :

```shell
yarn global add serve
server -s .
```

For more information on how the plugin loading works, visit this repository : [polymer-application-loader](https://github.com/Sehsyha/polymer-application-loader)