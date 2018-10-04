
var fs = require('fs');
var path = require('path');
var args = process.argv.splice(2);
var command = args.shift();

var taskDescription = args.join(' ');
var file = path.join(process.cwd(), '/.tasks');

switch (command) {
        case 'list':
                listTasks(file);
                break;
        case 'add':
                addTask(file, taskDescription);
                break;
        default:
                console.log('only list|add supported');
}

function loadOrInitializedTaskArray(file, cb) {
        fs.exists(file, function(exists) {
                var tasks = [];
                if (exists) {
                        fs.readFile(file, 'utf8', function(err, data) {
                                if (err) throw err;
                                var data = data.toString();
                                tasks = JSON.parse(data || '[]');
                                cb(tasks);
                        });
                } else {
                        cb([]);
                }
        });
}

function listTasks(file) {
        loadOrInitializedTaskArray(file, function(tasks) {
                for (var i in tasks) {
                        console.log(tasks[i]);
                }
        });
}

function storeTasks(file, tasks) {
        fs.writeFile(file, JSON.stringify(tasks), 'utf8', function(err) {
                if (err) throw err;
                console.log('Saved.');
        });
}

function addTask(file, taskDescription) {
        loadOrInitializedTaskArray(file, function(tasks) {
                tasks.push(taskDescription);
                storeTasks(file, tasks);
        });
}


