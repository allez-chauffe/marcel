"use strict";

const express = require('express');
const app = express();
const record = require('node-record-lpcm16')
const snowboy = require('snowboy');
const Models = snowboy.Models;
const Detector = snowboy.Detector;
const http = require('http').Server(app);
const ApiAi = require('apiai-promise')
const config = require('./config.js');
const apiai = ApiAi(config.apiaitoken);
const sessionId = 'marcel';
const models = new Models();

models.add({
  file: 'resources/marcel.pmdl',
  sensitivity: '0.4',
  hotwords: 'marcel'
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
  socket.on('speech', (speech) => {
    apiai.textRequest(speech.message, { sessionId })
         .then((response) => {
            console.log(response)
            if (response.result.metadata.intentName === 'YoutubeSearch') {
              io.sockets.emit('youtube', {type: "search", content: response.result.parameters.video});
            }

            if (response.result.metadata.intentName === 'YoutubePause') {
              io.sockets.emit('youtube', {type: "pause"});
            }

            if (response.result.metadata.intentName === 'YoutubePlay') {
              io.sockets.emit('youtube', {type: "play"});
            }

            if (response.result.metadata.intentName === "Planning") {
              io.sockets.emit('devfest', {type: "planning"});
            }

            if (response.result.metadata.intentName === 'CurrentTalk') {
              io.sockets.emit('devfest', {type: "current", location: response.result.parameters.location});
            }

            if (response.result.metadata.intentName === 'NextTalk') {
              io.sockets.emit('devfest', {type: "next", location: response.result.parameters.location});
            }

            if (response.result.metadata.intentName === 'CloseDisplayed') {
              io.sockets.emit('close');
            }

            if (response.result.metadata.intentName === 'Speaker') {
              io.sockets.emit('devfest', {type: "speaker", name: response.result.parameters.speaker});
            }

            if (response.result.metadata.intentName === 'Talk') {
              io.sockets.emit('devfest', {type: "talk", title: response.result.parameters.title});
            }
            
            if (response.result.metadata.intentName === 'Default Fallback Intent') {
              io.sockets.emit('default', response.result.fulfillment);
            }
          })
          .catch(err => console.log(err));
  });

})

detector.on('hotword', (index, hotword) => {
  console.log('hotword detected');
  io.sockets.emit('hotword');
})

const mic = record.start(config.microphone);
mic.pipe(detector);
