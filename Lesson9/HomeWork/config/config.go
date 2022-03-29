package config

import (
    "encoding/json"
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

type jsonConfig struct {
    Port        string `json:"port"`
    DbUrl       string `json:"db_url"`
    JaegerUrl   string `json:"jaeger_url"`
    SentryUrl   string `json:"sentry_url"`
    KafkaBroker string `json:"kafka_broker"`
    SomeAppId   string `json:"some_app_id"`
    SomeAppKey  string `json:"some_app_key"`
}

func FromFlags() (Config, error) {
    port := flag.String("port", "", "Номер порта")
    dbUrl := flag.String("db-url", "", "URL подключения к БД")
    jaegerUrl := flag.String("jaeger-url", "", "URL подключения к Jaeger")
    sentryUrl := flag.String("sentry-url", "", "URL подключения к Sentry")
    kafkaBroker := flag.String("kafka-broker", "", "Сервис Kafka")
    someAppId := flag.String("some-app-id", "", "Some app id")
    someAppKey := flag.String("some-app-key", "", "Some app key")

    flag.Parse()

    return FromRaw(*port, *dbUrl, *jaegerUrl, *sentryUrl, *kafkaBroker, *someAppId, *someAppKey)
}

func FromEnv() (Config, error) {
    port := os.Getenv("PORT")
    dbUrl := os.Getenv("DB_URL")
    jaegerUrl := os.Getenv("JAEGER_URL")
    sentryUrl := os.Getenv("SENTRY_URL")
    kafkaBroker := os.Getenv("KAFKA_BROKER")
    someAppId := os.Getenv("SOME_APP_ID")
    someAppKey := os.Getenv("SOME_APP_KEY")

    return FromRaw(port, dbUrl, jaegerUrl, sentryUrl, kafkaBroker, someAppId, someAppKey)
}

func FromRaw(port, dbUrl, jaegerUrl, sentryUrl, kafkaBroker, someAppId, someAppKey string) (Config, error) {
    if port == "" {
        return Config{}, errors.New("Необходимо указать порт")
    }
    portInt, err := strconv.Atoi(port)
    if err != nil {
        return Config{}, err
    }

    dbUrlStruct, err := parseUrl(dbUrl, "db-url")
    if err != nil {
        return Config{}, err
    }
    jaegerUrlStruct, err := parseUrl(jaegerUrl, "jaeger-url")
    if err != nil {
        return Config{}, err
    }
    sentryUrlStruct, err := parseUrl(sentryUrl, "sentry-url")
    if err != nil {
        return Config{}, err
    }

    return Config{
        Port:        portInt,
        DbUrl:       dbUrlStruct,
        JaegerUrl:   jaegerUrlStruct,
        SentryUrl:   sentryUrlStruct,
        KafkaBroker: kafkaBroker,
        SomeAppId:   someAppId,
        SomeAppKey:  someAppKey,
    }, nil
}

func FromJsonFile(pathToFile string) (Config, error) {
    content, err := os.ReadFile(pathToFile)
    if err != nil {
        return Config{}, err
    }

    config := jsonConfig{}
    err = json.Unmarshal(content, &config)
    if err != nil {
        return Config{}, err
    }

    return FromRaw(config.Port, config.DbUrl, config.JaegerUrl, config.SentryUrl, config.KafkaBroker, config.SomeAppId, config.SomeAppKey)
}

func parseUrl(raw, param string) (*url.URL, error) {
    res, err := url.ParseRequestURI(raw)
    if err != nil {
        return nil, errors.New("Неверный формат URL для параметра " + param)
    }

    return res, nil
}
