package config

import (
	"errors"
	"flag"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DbUrl       *url.URL
	JaegerUrl   *url.URL
	SentryUrl   *url.URL
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

func FromFlags() (Config, error) {
	port := flag.Int("port", 0, "Номер порта")
	dbUrlRaw := flag.String("db-url", "", "URL подключения к БД")
	jaegerUrlRaw := flag.String("jaeger-url", "", "URL подключения к Jaeger")
	sentryUrlRaw := flag.String("sentry-url", "", "URL подключения к Sentry")
	kafkaBroker := flag.String("kafka-broker", "", "Сервис Kafka")
	someAppId := flag.String("some-app-id", "", "Some app id")
	someAppKey := flag.String("some-app-key", "", "Some app key")

	flag.Parse()

	dbUrl, err := parseUrl(*dbUrlRaw, "db-url")
	if err != nil {
		return Config{}, err
	}
	jaegerUrl, err := parseUrl(*jaegerUrlRaw, "jaeger-url")
	if err != nil {
		return Config{}, err
	}
	sentryUrl, err := parseUrl(*sentryUrlRaw, "sentry-url")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:        *port,
		DbUrl:       dbUrl,
		JaegerUrl:   jaegerUrl,
		SentryUrl:   sentryUrl,
		KafkaBroker: *kafkaBroker,
		SomeAppId:   *someAppId,
		SomeAppKey:  *someAppKey,
	}, nil
}

func FromEnv() (Config, error) {
	portRaw := os.Getenv("PORT")
	dbUrlRaw := os.Getenv("DB_URL")
	jaegerUrlRaw := os.Getenv("JAEGER_URL")
	sentryUrlRaw := os.Getenv("SENTRY_URL")
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	someAppId := os.Getenv("SOME_APP_ID")
	someAppKey := os.Getenv("SOME_APP_KEY")

	if portRaw == "" {
		return Config{}, errors.New("Необходимо указать порт")
	}
	port, err := strconv.Atoi(portRaw)
	if err != nil {
		return Config{}, err
	}

	dbUrl, err := parseUrl(dbUrlRaw, "db-url")
	if err != nil {
		return Config{}, err
	}
	jaegerUrl, err := parseUrl(jaegerUrlRaw, "jaeger-url")
	if err != nil {
		return Config{}, err
	}
	sentryUrl, err := parseUrl(sentryUrlRaw, "sentry-url")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:        port,
		DbUrl:       dbUrl,
		JaegerUrl:   jaegerUrl,
		SentryUrl:   sentryUrl,
		KafkaBroker: kafkaBroker,
		SomeAppId:   someAppId,
		SomeAppKey:  someAppKey,
	}, nil
}

func parseUrl(raw, param string) (*url.URL, error) {
	res, err := url.ParseRequestURI(raw)
	if err != nil {
		return nil, errors.New("Неверный формат URL для параметра " + param)
	}

	return res, nil
}
