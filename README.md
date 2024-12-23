# Веб-сервис для решения простейших арифметических выражений


## Вводимые данные


Сервер принимает запрос по url-ом `/api/v1/calculate` и телом  `{
    "expression": "выражение, которое ввёл пользователь"
}`


Сервер поддерживает следующие арифметические операции с числами (именно числами, т.е. имеющими более 1 цифры в записи):  
* умножение
* деление
* сложение
* вычитание


А также учитывает приоритетные знаки (в том числе и скобки)


Наличие и количество пробелов в вводимом выражении неважно


## Ошибки


Сервер будет выдавать ошибки в следующих случаях:  
* Введенное выражение имеет незакрытые скобки или несколько операторов, стоящих подряд ("invalid expression")
* В результате вычислений будет происходить деление на ноль ("division by zero")
* Введенное выражение имеет символы кроме символов операторов `()+-.*` и цифр `1234567890` ("invalid expression")
* Была введена пустая строка ("invalid body")
* При попытке отправить запрос не POST методом ("wrong method")


## Запуск проекта

Для начала установите Go на свой компьютер [install golang](https://go.dev/doc/install), VisualStudio[install VS](https://code.visualstudio.com/)

Далее, зайдя в VStudio необходимо клонировать репозиторий `git clone https://github.com/h3xhmmr/calc_go`, а затем запустить сервер (по умолчанию работает на порте :8080) `go run calc_go/cmd/main.go`


## Примеры использования

Код ошибки 200  
`curl --location http://localhost:8080/api/v1/calculate --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`  
Результат:  
`"result":"6.00"`

Код ошибки 422  
`curl --location http://localhost:8080/api/v1/calculate --header 'Content-Type: application/json' --data '{"expression": "1/0"}'`  
Результат:  
`"error":"division by zero"`  

Код ошибки 400
`curl --location http://localhost:8080/api/v1/calculate --header 'Content-Type: application/json' --data '{"expression": ""}'`  
Результат:  
`"error":"invalid body"`  

Также возможен (желательно, т.к. curl немного сломан) запуск через Postman  
Для этого в строке адреса выберите метод Post и введите адрес `http://localhost:8080/api/v1/calculate`, после чего в поле body введите свой запрос в формате `{"expression": "2+2*2"}`


Примеры использования через Postman находятся в папке example_postman
