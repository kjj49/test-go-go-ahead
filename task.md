Тестовое задание на позицию Golang Developer

Задача:

Сделать API сервис который по запросу ([GET] /currency) возвращает курс валюты в виде json.

Сервис должен принимать query параметры:

date — дата для получения котировок на заданный день (по умолчанию текущий день);

val — валюта, по которой хотим получить курс.
  
Ответ сервиса:

В ответе должен быть курс валюты за указанный день.

Для получения данных курса можно использовать любые API, например:

https://cbr.ru/development/SXML/

Сервис должен быть обернут в Docker и опубликован на Github.

Будет большим плюсом:

Использование чистой архитектуры в проекте;

Покрытие сервиса тестами (e2e/unit);

Документация по проекту (как запустить локально);

Работа сервиса без участия пользователя (ежедневно в 10:00 UTC+ собирать котировки без запроса пользователя) и сохранять их в базе данных.

Решение просьба отправить ссылкой на Github-репозиторий.
