"use strict";

const express = require('express');
const app = express();
const record = require('node-record-lpcm16')
const snowboy = require('snowboy');
const Models = snowboy.Models;
const Detector = snowboy.Detector;
const http = require('http').Server(app);

const models = new Models();

models.add({
    file: 'resources/snowboy.umdl',
    sensitivity: '0.8',
    hotwords: 'snowboy'
})

const detector = new Detector({
    resource: 'resources/common.res',
    audioGain: 2.0,
    models
})

app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'X-Requested-With')
    next()
})

app.use(express.static('./public/'));
app.use(express.static('./components/'));
app.use(express.static('./node_modules/'));

app.get('/', (req, res) => {
    res.sendfile("public/index.html");
});

/**
 * Server itself
 * @type {http.Server}
 */
const server = app.listen(8080, () => {
    //print few information about the server
    const host = server.address().address;
    const port = server.address().port;
    console.log("Server running and listening @ " + host + ":" + port);
});
const io = require('socket.io')(server);

io.on('connection', socket => {
    socket.on('speech', (message) => console.log(message))
})

detector.on('hotword', (index, hotword) => {
    console.log('hotword detected');
    io.sockets.emit('hotword');
})

/** list of components to be loaded */
const componentsList = {
    "styles": [
        "css/style.css",
        "css/font-awesome.min.css",
        "css/weather-icons.min.css"
    ],
    "scripts": [
        "socket.io/socket.io.js",
        "js/connect-socketio.js"
    ],
    "components": [
        {
            "componentName": "youtube",
            "eltName": "youtube-item",
            "files": "youtube.html"
        },
        {
            "componentName": "marcel",
            "eltName": "marcel-item",
            "files": "marcel.html",
            "propValues": {
                "logo_url": "assets/zenika.png",
                "message_text1": "Bienvenue",
                "message_text2": "à Zenika Lille",
                "github_users": [
                    'Gillespie59',
                    'GwennaelBuchet',
                    'T3kstiil3',
                    'RemiEven',
                    'looztra',
                    'a-cordier',
                    'wadendo',
                    'NathanDM',
                    'Antoinephi',
                    'cluster',
                    'yyekhlef',
                    'gdrouet',
                    'Kize',
                    'kratisto',
                    'Sehsyha',
                    'P0ppoff'
                ],
                "github_client_id": "27b62f039b44ddc08fdf",
                "github_client_secret": "7b14f465112e87267a72c02d4c3fc58925412dbd",
                "twitter_api": "http://10.0.10.63:8090/api/v1/twitter/timeline",
                "vlille_stations_id": [
                    { name: "Rihour", id: 10},
                    { name: "Cormontaigne", id: 36 },
                    { name: "Mairie de Lille", id: 64 },
                    { name: "Gare Lille Flandres", id: 25 },
                    { name: "Boulevard Louis XIV", id: 47 }
                ],
                "soundtouch_url": "http://10.0.10.166:8090/now_playing",
                "weather_api_key": "FREE_OPENWEATHER_KEY",
                "weather_city": "Lille,Fr",
                "weather_url": "http://10.0.10.63:8090/api/v1/weather/forecast/5",
                "calendar_url": "http://10.0.10.63:8090/api/v1/agenda/incoming/50?json_callback=JSON_CALLBACK",
                "speech_default_message": "Bonjour à tous, je suis MARCEL !"
            }
        }
    ]
};

/**
 * Get a list of JSON for all registered components
 * @path /componentsList
 * @HTTPMethod GET
 * @returns {string}
 */
app.get("/componentsList", function (req, res) {
    res.send(componentsList);
});
