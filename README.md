# skvdmt/email
Golang package for sending email from smtp server.

Can use auth on smtp server and connect without authorization.

## Instalation
`go get https://github.com/skvdmt/email`

## Usage

make email:
```go
var let email.Letter
let.Set("to", "email to")
let.Set("content-type", "text/plain")
let.Set("subject", "test subject")
let.Set("body", "test body")
```
set local server and send email:
```go
var serv email.Server
serv.Set("host", "localhost")
serv.Set("port", 25)
conn := serv.Conn()
email.Send(&serv, &let, conn)`
```
or set smtp server with auth and send email: 
```go
serv.Set("host", "smtp.yandex.ru")
serv.Set("port", 587)
serv.Set("username", "your_login@rambler.ru")
serv.Set("password", "your_password")
conn := serv.Auth()
email.Send(&serv, &let, conn)`
```
