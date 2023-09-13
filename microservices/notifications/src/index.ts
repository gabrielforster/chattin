import WebSocket from "ws";

const connections: Array<WebSocket> = [];

const server = new WebSocket.Server({ port: 42691 });

server.on("connection", (socket) => {
  connections.push(socket);

  socket.on("message", (message) => {
    const messageAsJSON = JSON.parse(message.toString());

    connections.forEach((connection) => {
      connection.send(Buffer.from(messageAsJSON.message));
    });
  });
});
