import WebSocket from "ws";
import { createKafkaConsumer } from "./lib/kafka";

const MS_NOTIFICATION_WEBSOCKET_PORT = Number(process.env.MS_NOTIFICATION_WEBSOCKET_PORT) || 42691;
const MS_NOTIFICATION_PORT = process.env.MS_NOTIFICATION_PORT || 42692;
const MESSAGE_KAFKA_TOPIC = "messages"

const connections: Array<WebSocket> = [];

const server = new WebSocket.Server({ port: MS_NOTIFICATION_WEBSOCKET_PORT });

server.on("connection", (socket) => {
  connections.push(socket);

  socket.on("message", (message) => {
    const messageAsJSON = JSON.parse(message.toString());

    connections.forEach((connection) => {
      connection.send(Buffer.from(messageAsJSON.message));
    });
  });
});

(async () => {
  const consumer = createKafkaConsumer({ topic: MESSAGE_KAFKA_TOPIC });
  consumer.on("message", (message) => {
    connections.forEach((connection) => {
      connection.send(message.value);
    });
  })
})()
