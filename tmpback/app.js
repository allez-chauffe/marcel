"use strict";

let express = require('express');
let app = express();
let http = require('http').Server(app);

app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'X-Requested-With')
    next()
})

app.use(express.static('./public/'));
app.use(express.static('./components/'));
app.use(express.static('./node_modules/'));



app.get('/', function (req, res) {
    res.sendfile("public/index.html");
});

/**
 * Server itself
 * @type {http.Server}
 */
let server = app.listen(8080, function () {
    //print few information about the server
    let host = server.address().address;
    let port = server.address().port;
    console.log("Server running and listening @ " + host + ":" + port);
});

/** list of components to be loaded */
let componentsList = [
    {
        "componentName": "weather",
        "eltName": "weather-item",
        "files": "weather.html",
        "propValues": {
            "city": "Lille,Fr",
            "weather_url": "http://10.0.10.63:8090/api/v1/weather/forecast/5?",
            "api_key": "OPENWEATHER_FREE_API_KEY"
        }
    }
];

/**
 * Get a list of JSON for all registered components
 * @path /componentsList
 * @HTTPMethod GET
 * @returns {string}
 */
app.get("/componentsList", function (req, res) {
    res.send(componentsList);
});
