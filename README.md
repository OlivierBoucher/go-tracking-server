# go-tracking-server
A draft for an event tracking server in Go  
  [![Build Status](https://travis-ci.org/OlivierBoucher/go-tracking-server.svg?branch=master)](https://travis-ci.org/OlivierBoucher/go-tracking-server)
[![Coverage Status](https://coveralls.io/repos/OlivierBoucher/go-tracking-server/badge.svg?branch=master&service=github)](https://coveralls.io/github/OlivierBoucher/go-tracking-server?branch=master)

##TODO
- [ ] Make package names more relevant, merge/decouple packages that require so
- [ ] Implement file logging for production + rotation
- [ ] More tests
- [x] Complete TODOS concerning logging
- [x] Complete TODOS concerning error handling - Decide actions to undertake based on situations
- [x] Implement a protocol over websocket (Connection, Disconnection, Validation etc)
- [x] File based configuration for connections (DB, AMQP) in JSON - VERY BASIC


##First pitch
![Imgur](http://i.imgur.com/rqwf7Yc.png)
