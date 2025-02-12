
var events = require('events');
var net = require('net');
var channel = new events.EventEmitter();
channel.clients = {};
channel.subscriptions = {};

channel.on('join', function(id, client) {
        this.clients[id] = client;
        this.subscriptions[id] = function(senderId, msg) {
                if (id != senderId) {
                        this.clients[id].write(msg);
                }
        }
        this.on('broadcast', this.subscriptions[id]);
});

channel.on('leave', function(id) {
        channel.removeListener('broadcast', this.subscriptions[id]);
        channel.emit('broadcast', id, id + ' has left the chat.\n');
});

var server = net.createServer(function(client) {
        var id = client.remoteAddress + ':' + client.remotePort;
        channel.emit('join', id, client);
        client.on('data', function(data) {
                var msg = id + ' > ' + data.toString();
                channel.emit('broadcast', id, msg);
        });
        client.on('close', function() {
                channel.emit('leave', id);
        });
});

server.listen(8888);

