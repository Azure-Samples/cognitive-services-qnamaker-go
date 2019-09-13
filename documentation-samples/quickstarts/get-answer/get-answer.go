package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
)

/*
1. Create environment variabes QNA_MAKER_RESOURCE_ENDPOINT, QNA_MAKER_ENDPOINT_KEY, QNA_MAKER_KB_ID.
2. Compile with: go build get-answer.go
3. Execute with: ./get-answer
*/

func main() {

    // Your QnA Maker resource endpoint.
    // From Publish Page: HOST
    // Example: https://YOUR-RESOURCE-NAME.azurewebsites.net/
    if "" == os.Getenv("QNA_MAKER_RESOURCE_ENDPOINT") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_RESOURCE_ENDPOINT.")
    }
    var endpoint string = os.Getenv("QNA_MAKER_RESOURCE_ENDPOINT")

    // Authorization endpoint key
    // From Publish Page
    // Note this is not the same as your QnA Maker subscription key.
    if "" == os.Getenv("QNA_MAKER_ENDPOINT_KEY") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_ENDPOINT_KEY.")
    }
    var endpoint_key string = os.Getenv("QNA_MAKER_ENDPOINT_KEY")

    // QnA Maker Knowledge Base ID
    if "" == os.Getenv("QNA_MAKER_KB_ID") {
        log.Fatal("Please set/export the environment variable QNA_MAKER_KB_ID.")
    }
    var kb_id string = os.Getenv("QNA_MAKER_KB_ID")

    // Management APIs postpend the version to the route
    // From Publish Page, value after POST
    // Example: /knowledgebases/ZZZ15f8c-d01b-4698-a2de-85b0dbf3358c/generateAnswer
    var route string = "/qnamaker/knowledgebases/" + kb_id + "/generateAnswer";

    // JSON format for passing question to service
    var question string = "{'question': 'Is the QnA Maker Service free?','top': 3}"

    req, _ := http.NewRequest("POST", endpoint + route, bytes.NewBuffer([]byte(question)))
    req.Header.Add("Authorization", "EndpointKey " + endpoint_key)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Content-Length", strconv.Itoa(len(question)))
    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)

    fmt.Printf(string(body) + "\n")
}
