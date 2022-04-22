package main

import (
	"GeekBrains/Lesson8/HomeWork/config"
	"fmt"
)

func main() {
	// go run main.go
	//configRaw, err := config.FromRaw("8080", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "http://jaeger:16686", "http://sentry:9000", "kafka:9092", "testid", "testkey")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(configRaw)

	// go run main.go --port=8080 --db-url="postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" --jaeger-url=http://jaeger:16686 --sentry-url=http://sentry:9000 --kafka-broker=kafka:9092 --some-app-id=testid --some-app-key=testkey
	//configFlags, err := config.FromFlags()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(configFlags)

	// PORT=8080 DB_URL="postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" JAEGER_URL=http://jaeger:16686 SENTRY_URL=http://sentry:9000 KAFKA_BROKER=kafka:9092 SOME_APP_ID=testid SOME_APP_KEY=testkey go run main.go
	configEnv, err := config.FromEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println(configEnv)
}
