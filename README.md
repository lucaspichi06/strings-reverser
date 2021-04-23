# String Reverser ![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This API is responsible for provide an endpoint to reverse the words in a given sentence.

For example:
````
input: today is the first day of the rest of my life
output: life my of rest the of day first the is today
````

## Index
- [Run it locally](#run-it-locally)
- [Run it from the cloud](#run-it-from-the-cloud)
- [Test the application](#test-the-application)
- [Solution approach](#solution-approach)

## Run it locally
To run the api locally it is necessary to clone this repository from Github:
````
git clone https://github.com/lucaspichi06/strings-reverter.git
````

After that, you have to move to the root of the project and run this command:
````
go run src/api/main.go
````

Now you have the API running in the port :8080 so you can invoke the endpoint with the following curl (you can use Postman too):
````bash
curl --location --request POST 'http://localhost:8080/revert_string' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "today is the first day of the rest of my life"
}'
`````

It will respond in the following way:
````JSON
{
    "message_to_revert": "today is the first day of the rest of my life",
    "reverted_message": "life my of rest the of day first the is today"
}
````

In case of error it will respond with the right status code and the following response:
````JSON
{
    "message": "error while reverting the message",
    "error": "internal_server_error",
    "status": 500,
    "cause": [
        "ups... there is nothing to revert"
    ]
}
````

Besides that, you have another endpoint to check the API health status:
````bash
curl --location --request GET 'http://localhost:8080/ping'
````

If the API is running successfully, it has to return the word ```pong``` 

## Run it from the cloud
If you don't want to clone the repository and you need to test the API you can to call this application using the following curl  (you can use Postman too):
````bash
curl --location --request POST 'https://strings-reverter.herokuapp.com/revert_string' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": ""
}'
````

And to see the API health status:
````bash
curl --location --request GET 'http://localhost:8080/ping'
````

#### **** You can find the postman collection in the utils folder! ****

## Test the application
This API have unit tests to warranties the integrity of the application.
You can run this tests with the following command from the root of the project:
````
go test ./...
````

#### Coverage (85.7%)
You can also generate a coverage report with the following commands:
- Generating the coverage.out file
````
go test -covermode=set -coverprofile=coverage.out ./...
````
- Displaying an HTML coverage report
 
````
 go tool cover -html=coverage.out
````

## Solution approach
The approach I followed was creating an API who gives an endpoint to reverse an input sentence. 
This endpoint is a post and receives in the body of the request the message to reverse.

For the API architecture I used an Hexagonal Architecture (maybe is too much for this problem) and implemented the inversion of control through dependency injection.
![Hexagonal Architecture](https://miro.medium.com/max/1718/1*yR4C1B-YfMh5zqpbHzTyag.png)

The Revert function is the core of the application.
It takes an request struct who has the message to reverse as input.
Then It splits the message into an slice and iterate the slice to build the response.