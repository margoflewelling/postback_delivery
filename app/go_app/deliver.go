// pull from redis queue
// deliver get response
// log delivery time, response code, response time, response body

package main

import (
        "context"
        "fmt"
        "time"
        "github.com/go-redis/redis"
        "encoding/json"
        "strings"
        "net/http"
      )
var ctx = context.Background()

// set up object to hold request information that is coming in as hash
type Request struct {
  Endpoint struct {
                    Method string `json: "method"`
                    Url string `json: "url"`
                   } `json: "endpoint"`
  Data []struct {
                  Mascot string `json: "mascot"`
                  Location string `json: "location"`

  } `json:"data"`
}
// set up redis client
func main() {
  fmt.Println("Starting")
    RedisClient := redis.NewClient(&redis.Options{
      Addr: "redis:6379",
    })
    // subscribe to the channel 'taskQ' established in php app
    subNewMessage := RedisClient.Subscribe(ctx, "taskQ")
    ch := subNewMessage.Channel()
    // time out after 60 sec of listening
    time.AfterFunc(time.Second * 60, func(){
      _ = subNewMessage.Close()
    })
    // send responses for each incoming message from the queue
    for msg := range ch {
      var request Request
      // fmt.Println(msg.Channel, msg.Payload)
      // put msg body in the request struct
      err := json.Unmarshal([]byte(msg.Payload), &request)
      if err != nil {
        fmt.Println("error:", err)
      }
      url := request.Endpoint.Url
      // should make dynamic with data keys
      endpoint := strings.Replace(url, "{mascot}", request.Data[0].Mascot, 1)
      endpoint = strings.Replace(endpoint, "{location}", request.Data[0].Location, 1)
      send_request(request.Endpoint.Method, endpoint)
    }
    fmt.Println("Stopping")
}

// method to send http response
func send_request(method string, url string) {
    // make http get request
    switch method {
      case "GET":
        start := time.Now()
        resp, err := http.Get(url)
        if err != nil {
          fmt.Println("error:", err)
        }
        // log data to console
        fmt.Println("Response code:", resp.StatusCode)
        fmt.Println("Delivery time:", start)
        fmt.Println("Response Time:", time.Now())
        fmt.Println("Response Body:", resp.Body)
      }
}
