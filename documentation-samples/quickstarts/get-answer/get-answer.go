package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
)

// **********************************************
// *** Update or verify the following values. ***
// **********************************************

// Represents the various elements used to create HTTP request URIs
// for QnA Maker operations.
// From Publish Page: HOST
// Example: https://YOUR-RESOURCE-NAME.azurewebsites.net/qnamaker
string host = "https://YOUR-RESOURCE-NAME.azurewebsites.net/qnamaker";

// Authorization endpoint key
// From Publish Page
string endpoint_key = "YOUR-ENDPOINT-KEY";

// Management APIs postpend the version to the route
// From Publish Page, value after POST
// Example: /knowledgebases/ZZZ15f8c-d01b-4698-a2de-85b0dbf3358c/generateAnswer
string route = "/knowledgebases/YOUR-KNOWLEDGE-BASE-ID/generateAnswer";

// JSON format for passing question to service
string question = @"{'question': 'Is the QnA Maker Service free?','top': 3}";

func main() {

	req, _ := http.NewRequest("POST", host + route, bytes.NewBuffer([]byte(question)))
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

    fmt.Printf(body + "\n")
}