# shortener

## О проекте

Проект выполнен по заданию Ozon Fintech и является сервисом для укорачивания ссылок с использованием grpc-gateway. Сервис запускается с помощью docker-compose. В качестве хранилища используется PostgreSQL или in-memory, выбор происходит при запуске.
## Запросы к сервису
 - ```GET /get/{short_link}```
 - ```POST /create body: {"link": "https://google.com/"}```
 ### Примеры запросов
 ```shell
 # POST
 > curl -X POST localhost:8090/create -H "Content-Type: application/json" -d '{"link": "google.com"}'
 # GET
 > curl localhost:8090/get/huPVAz7b6R
 ```
## Запуск
Для запуска необходимо выполнить следующие команды:
```shell
# Клонируем репозиторий
> git clone https://github.com/FastPretzel/shortener
# Запуск кодогенерации из proto файла
> make generate-grpc-gateway
# Запуск в режиме postgresql
> make postgresql
# Запуск в режиме in-memory
> make memory
# Запуск тестов
> make test
```
