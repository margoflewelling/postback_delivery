# postback_delivery

### Data Distribution Simulation
- Php application to handle incoming http post requests /app/php
- Go application to deliver http responses /app/go
- Redis queue between the php and go applications


### To run locally:
The php app:
- redis-server
- redis-cli monitor  (to see subscriptions and objects published to the redis channel)
- php -S localhost:8000  (run the ingestion agent on port 8000 )
The go app(dependent on the php + redis):
- go run app/go_app/deliver.go (should see starting printed to the console,
               and then a log of responses in response to postman request.
                Will time out after 1 minute)


- tested locally using postman
#### Post to http://localhost:8000 with this body format:
```
    {
    "endpoint":{
    "method":"GET",
    "url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"
    },
    "data":[
    {
    "mascot":"Gopher",
    "location":"https://blog.golang.org/gopher/gopher.png"
    }
    ]
    }
```

### Dockerizing
  Run docker-compose up
