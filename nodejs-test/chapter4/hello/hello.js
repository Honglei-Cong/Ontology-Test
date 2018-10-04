
var http = require('http');
var server = http.createServer(function(req, res) {
        var body = 'Hello';
        res.setHeader('Content-Length', body.length);
        res.setHeader('Content-Type', 'Text/plain');
        res.end(body);
});
server.listen(8888);

