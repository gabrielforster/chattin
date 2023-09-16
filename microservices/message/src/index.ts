import { v4 as uuidv4 } from "uuid"
import express from "express"
import { createKafkaProducer } from "./lib/kafka"

const MS_MESSAGE_PORT = process.env.MS_MESSAGE_PORT ? Number(process.env.MS_MESSAGE_PORT) : 42693;
const KAFKA_MESSAGE_TOPIC = "messages";

(async () => {
  const app = express()
  app.use(express.json())
  const kafkaProducer = await createKafkaProducer()

  app.get("/health/check", (_, res) => {
    res.status(200).json({ message: "message microservice health check ok" })
  })

  app.post("/message", (req, res) => {
    const { message: text } = req.body as { message: string }

    const message = {
      id: uuidv4(),
      text,
      status: "sent",
      createdAt: new Date().getTime(),
    }

    kafkaProducer.send(
      [{ topic: KAFKA_MESSAGE_TOPIC, messages: Buffer.from(JSON.stringify(message)) }],
      (err, data) => {
        if (err) {
          console.info(`Could not produce kafka message for message ${message.id}`)
        }

        console.info(`Produced kafka message for message ${message.id}`, data)
      })


    res.status(200).json({ message })
  })

  app.listen(MS_MESSAGE_PORT, () => {
    console.log(`message microservice listening on port ${MS_MESSAGE_PORT}`)
  })
})()
