# go-vuejs-chat
Simple Go Vue.js chat application

Source Code for the tutorial @whichdev.com:
https://www.whichdev.com/go-vuejs-chat/

## Login TEST
```
curl --location --request GET 'localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "john",
    "Password":"password"
}'
```