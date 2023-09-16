import { Producer, KafkaClient } from "kafka-node"

export function createKafkaProducer(kafkaHost?: string | undefined): Promise<Producer> {
  return new Promise((resolve, reject) => {
    const client = new KafkaClient({ kafkaHost: kafkaHost ?? "localhost:9092" })
    const producer = new Producer(client)

    producer.on("ready", () => {
      resolve(producer)
    })

    producer.on("error", (err) => {
      reject(err)
    })
  })
}
