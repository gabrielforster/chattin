import { KafkaClient, Consumer } from "kafka-node"

export function createKafkaConsumer({ kafkaHost, topic }: { kafkaHost?: string | undefined , topic: string  }): Consumer {
  const client = new KafkaClient({ kafkaHost: kafkaHost ?? "localhost:9092" })
  return new Consumer(client, [{ topic }], { autoCommit: false })
}
