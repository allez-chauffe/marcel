# Frontend

This part of the project is what will be on the mirror. It's a lightweight html application with the inclusion of Polymer. At startup, it will load the list of plugins which have to be displayed, than load every plugin into the page and finally display them.

## Run

You can run the application with docker by running the commands in this directory :

```shell
docker build -t marcel-front .
docker run -p 80:80 -v /path/to/plugins/directory:/usr/share/nginx/html/plugins -it marcel-front
```

Note that you also have to run the backend to have the list of plugins.

## Installation

First, you have to fetch the bower components:

```bash
bower install
```

Then you have to copy the config.example.json file into config.json and change the addesses to specify the URL of the plugin list and the plugins.

Finally you just have to server the files with your favorite server like nginx or with :

```shell
yarn global add serve
server -s .
```

For more information on how the plugin loading works, visit this repository : [polymer-application-loader](https://github.com/Sehsyha/polymer-application-loader)