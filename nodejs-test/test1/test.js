var fs = require('fs');
fs.readFile("./a.json", function(er, data) {
    console.log(data);
})