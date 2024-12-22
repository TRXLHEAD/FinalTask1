### Проект калькулятор на http сервере go

# Калькулятор содержит функции:
> Сложение <
> Вычитание <
> Умножение <
> Деление <

# Для запуска на Visual Studio Code (Windows):
> Сохраните проект (обязательно извлеките папку Calc из FinalTask1, иначе вывдаст ошибку)
> Откройте терминал, напишите: go run ./cmd/main.go
  Если появилось сообщение: "Сервер успешно запущен", значит сервер работает корректно
> После запуска сервера используйте второй терминал: Alt + Shift + F5 (Если не открывается, можно нажать на '+' в терминале)
> Введите команду: Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
> Готово

# Пример использования:
> (5+6) / 2
> 5,5
----------------
> 2+2*4
> 10

# Статусы сервера:

### ✅ 200 - StatusOK > Корректная работа сервера, ответ получен

### ❌ 405 - StatusMethodNotAllowed > Метод, используемый в запросе, не соответствует для данного сервера. в прошлом important сказано, какой стоит использовать. убедитесь, что используете метод POST

### ❌ 500 - StatusInternalServerError > Ошибка, что-то пошло не так

# Пример ответов сервера:

### ✅ 200 (StatusOK)

```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```
### результат
```
{"result":"6.00"}
```
### ❌ 422 (StatusUnprocessableEntity) ошибка пользователя в написании выражения
```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*"}' | Select-Object -Expand Content
```
### результат
```
Invoke-WebRequest : {"error":"expression is not valid"}
строка:1 знак:1
+ Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/c ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
### ❌ 405 (StatusMethodNotAllowed) не правильный метод
```
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```
### результат
```
Invoke-WebRequest : Невозможно отправить тело содержимого с данным типом предиката.
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/ca ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (:) [Invoke-WebRequest], ProtocolViolationException
    + FullyQualifiedErrorId : System.Net.ProtocolViolationException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
