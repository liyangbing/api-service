package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "http://20.75.203.79:50002/chat"
   method := "POST"

   payload := strings.NewReader(`{
    "key": "n9qCDwTD",
    "prompt": "你好，你是谁",
    "type": "text"
}`)

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
   req.Header.Add("Content-Type", "application/json")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}