# email-sender
my test email sender


**query example**
```sh
curl --location --request POST 'localhost:8080/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "reciever": "u.shylianok@gmail.com",
    "text": "Привет! Это мое тестовое собщение"
}'
```