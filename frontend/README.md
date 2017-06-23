# Frontend

This part of the project is what will be on the mirror. It's a lightweight html application with the inclusion of Polymer. At startup, it will load the list of plugins which have to be displayed, than load every plugin into the page and finally display them.

## Installation

First, you have to fetch the bower components:

```bash
bower install
```

Then you have to copy the config.example.json file into config.json and change the addesses to specify the location of the plugin list and the plugins.

Finally you just have the files with your favorite server like nginx.

For more information on how the plugin loading works, visit this repository : [polymer-application-loader](https://github.com/Sehsyha/polymer-application-loader)