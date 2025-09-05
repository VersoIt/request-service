package kafka

func (k *Kafka) Close() {
	k.client.Close()
}
