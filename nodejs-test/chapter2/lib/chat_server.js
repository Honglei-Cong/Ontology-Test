

function assignGuestName(socket, guestNumber, nickNames, namesUsed) {
    var name = 'Guest ' + guestNumber;
    nickNames[socket.id] = name;
    socket.emit('nameResult', {
        success: true,
        name: name
    });
    namesUsed.push(name);
    return guestNumber + 1;
}

function joinRoom(socket, room) {
    socket.join(room);
    currentRoom[socket.id] = room;
    socket.emit('joinResult', {room: room});
    socket.broadcast.to(room).emit('message', {
        text: nickNames[socket.id] + 'has joined ' + room + '.'
    });

    var usersInRoom = io.sockets.clients(romm);
    if (usersInRoom.length > 1) {
        var usersInRoomsSummary = "Users currently in " + room + ": ";
        for (var index in usersInRoom) {
            var userSocketId = usersInRoom[index].id;
            if (userSocketId != socket.id) {
                usersInRoomsSummary += ", ";
            }
            usersInRoomsSummary += nickNames[userSocketId];
        }
        usersInRoomsSummary += '.';
        socket.emit('message', {text: usersInRoomsSummary});
    }
}




exports.listen = function (server) {
  io = socketio.listen(server);
  io.set("log level", 1);
  io.socket.on('connection', function(socket) {
      guestNumber = assignGuestName(socket, guestNumber, nickNames, namesUsed);
      joinRoom(socket, "Lobby");
      handleMessageBroadcasting(socket, nickNames);
      handleNameChangeAttemps(socket, nickNames, namesUsed);
      handleRoomJoining(socket);

      socket.on('rooms', function() {
          socket.emit('rooms', io.sockets.manager.rooms);
      });
      handleClientDisconnection(socket, nickNames, namesUsed);
  })
};