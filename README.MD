# Тестовое задание для стажера-разработчика

## Сервис, предоставляющий API по созданию сокращённых ссылок.

Задача

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:
— Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;
— Длиной 10 символов;
— Из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Условия:
— Сервис распространён в виде Docker-образа;
— Хранилище in-memory и PostgreSQL. (параметр при запуске сервиса)
— Реализованный функционал покрыт Unit-тестами.

Сервис работает через GRPC.

# Работа сервиса

## Запуск
```bash
make inmemory

make postgresql

make test
```

## Отправка запросов

Я ипользовала для этого GUI "ezy": https://github.com/getezy/ezy

Пример:
```bash
127.0.0.1:8080

GetLink:
  Request:
  {
    "token": "D6T0pHJZlg"
  }
  Responce:
  {
  "url": "https://delinack.ru"
  }
  
CreateLink:
  Request:
  {
  "url": "https://delinack.ru"
  }
  Responce:
  {
  "token": "D6T0pHJZlg"
  }
```