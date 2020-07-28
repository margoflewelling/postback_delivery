<?php
// establish redis connection
$redis = new Redis();
// default redis port 6379
$redis->connect('redis', 6379);

// git the incoming http request method
$method = $_SERVER['REQUEST_METHOD'];

if ($method == "POST")
{
  echo "Pushing object to redis queue";
  // read incoming post body
  $json = file_get_contents('php://input');
  // decode to inspect in testing
  $request = json_decode($json, TRUE);
  // push object into redis channel named taskQ
  $redis -> publish("taskQ", json_encode($request));
}
else {
   echo "Only accepts POST requests";
  }

  // close down redis connection
$redis->close();

?>
