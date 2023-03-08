# notification

This is use for my little online class to teach how to use SMTP to send mail notif, and teach how to do an even drivien microservices as well


I use existing libs :

 - Chi Router
 - Ozzo Validation, for input request validation
 - Godotenv, for env loader



For setup :
- git clone 
- cd notification
- go mod tidy

The Endpoint :
```
curl --location --request POST 'http://localhost:8080/api/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sender":"jane.doe@gmail.com",
    "title":"no-reply",
    "message":"it just simple message",
    "recipients": [{
        "email":"jhon.doe@gmail.com"
    }]
```
}'