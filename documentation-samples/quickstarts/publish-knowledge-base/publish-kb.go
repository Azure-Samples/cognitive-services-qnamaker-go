package main

import (
	"bytes"
    "fmt"
    "net/http"
)

// 1. Replace variable values with your own from QnA Maker Publish page
// 2. Compile with: go build get-answer.go
// 3. Execute with: ./get-answer

func main() {

	var knowledge_base_id = "YOUR-KNOWLEDGE-BASE-ID";
	var resource_key = "YOUR-RESOURCE-KEY";

	var host = fmt.Sprintf("https://westus.api.cognitive.microsoft.com/qnamaker/v4.0/knowledgebases/%s", knowledge_base_id);
	var content = bytes.NewBuffer([]byte(nil));

	req, _ := http.NewRequest("POST", host, content)

	req.Header.Add("Ocp-Apim-Subscription-Key", resource_key)

    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }
}