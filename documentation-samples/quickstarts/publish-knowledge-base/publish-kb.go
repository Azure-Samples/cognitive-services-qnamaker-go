package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "os"
)

/*
1. Create environment variabes QNA_MAKER_ENDPOINT, QNA_MAKER_SUBSCRIPTION_KEY, QNA_MAKER_KB_ID.
2. Compile with: go build publish-kb.go
3. Execute with: ./publish-kb
4. For successful publish, no data is returned, only 204 http status
*/

func main() {

    // Your QnA Maker endpoint.
    var endpoint string = os.Getenv("QNA_MAKER_ENDPOINT")
    if "" == os.Getenv("QNA_MAKER_ENDPOINT") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_ENDPOINT.")
    }

    // QnA Maker subscription key
    // From Publish Page
    var subscription_key string = os.Getenv("QNA_MAKER_SUBSCRIPTION_KEY")
    if "" == os.Getenv("QNA_MAKER_SUBSCRIPTION_KEY") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_SUBSCRIPTION_KEY.")
    }

    // QnA Maker Knowledge Base ID
    if "" == os.Getenv("QNA_MAKER_KB_ID") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_KB_ID.")
    }
    var kb_id string = os.Getenv("QNA_MAKER_KB_ID")

    var service string = "/qnamaker/v4.0"
    var method string = "/knowledgebases/"
    var uri = endpoint + service + method + kb_id

    var content = bytes.NewBuffer([]byte(nil));

    req, _ := http.NewRequest("POST", uri, content)

    req.Header.Add("Ocp-Apim-Subscription-Key", subscription_key)

    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    // print 204 - success code
    fmt.Println(response.StatusCode)
}
