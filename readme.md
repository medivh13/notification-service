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


# update-1
in this update, I'll show about simple message broker that is Nats.
NATS is a simple, secure and high performance open source messaging system for cloud native applications, IoT messaging, and microservices architectures. The NATS server is written in the Go programming language, but client libraries to interact with the server are available for dozens of major programming languages. NATS supports both At Most Once and At Least Once delivery. It can run anywhere, from large servers and cloud instances, through edge gateways and even Internet of Things devices.

For a simple Nats's installation you can install nats into your device Via Docker, then set the Nats host in .env file.