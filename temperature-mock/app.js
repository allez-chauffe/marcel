"use strict";

const express = require('express');
const app = express();
const http = require('http').Server(app);

app.use((req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'X-Requested-With')
  next()
})

/**
 * Server itself
 * @type {http.Server}
 */
const server = app.listen(5000, () => {
  //print few information about the server
  const host = server.address().address;
  const port = server.address().port;
  console.log("Server running and listening @ " + host + ":" + port);
});
const io = require('socket.io')(server);

io.on('connection', socket => {
  setInterval(() => socket.emit('temperature', {name: 'Salle 1', value: '10'}), 3000)
})
