package config

import (
	"log"
	"os"
)

const (
	kafkaHost           = "KAFKA_HOST"
	schemaRegistryHost  = "SCHEMA_REGISTRY_HOST"
	BootstrapServers    = "bootstrap.servers"
	PearsonCreatedTopic = "PERSON_CREATED_EVENT"
	PersonSubjName      = "person.Event"
	ServerPort          = 9090
)

func GetKafkaHost() string {
	h := os.Getenv(kafkaHost)
	if h == "" {
		log.Panicf("%v must not be null", kafkaHost)
	}

	return h
}

func GetSchemaRegistryHost() string {
	h := os.Getenv(schemaRegistryHost)
	if h == "" {
		log.Panicf("%v must not be null", schemaRegistryHost)
	}

	return h
}
