package main

import (
	"fmt"
	"goLibrary/httpclient"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

func runHTTPClientExamples() {
	log := logger.NewLogger(logger.INFO)

	url := "https://jsonplaceholder.typicode.com/posts/1"
	body, err := httpclient.Get(url)
	if err != nil {
		log.Error("Failed to make GET request: " + err.Error())
	} else {
		log.Info("GET request successful: " + body)
		fmt.Println("GET Response:", body)
	}

	postUrl := "https://jsonplaceholder.typicode.com/posts"
	payload := `{"title":"foo","body":"bar","userId":1}`
	postBody, err := httpclient.Post(postUrl, "application/json", []byte(payload))
	if err != nil {
		log.Error("Failed to make POST request: " + err.Error())
	} else {
		log.Info("POST request successful: " + postBody)
		fmt.Println("POST Response:", postBody)
	}

	patchUrl := "https://jsonplaceholder.typicode.com/posts/1"
	patchPayload := `{"title":"updated title"}`
	patchBody, err := httpclient.Patch(patchUrl, "application/json", []byte(patchPayload))
	if err != nil {
		log.Error("Failed to make PATCH request: " + err.Error())
	} else {
		log.Info("PATCH request successful: " + patchBody)
		fmt.Println("PATCH Response:", patchBody)
	}

	putUrl := "https://jsonplaceholder.typicode.com/posts/1"
	putPayload := `{"id":1,"title":"updated title","body":"updated body","userId":1}`
	putBody, err := httpclient.Put(putUrl, "application/json", []byte(putPayload))
	if err != nil {
		log.Error("Failed to make PUT request: " + err.Error())
	} else {
		log.Info("PUT request successful: " + putBody)
		fmt.Println("PUT Response:", putBody)
	}

	deleteUrl := "https://jsonplaceholder.typicode.com/posts/1"
	deleteBody, err := httpclient.Delete(deleteUrl)
	if err != nil {
		log.Error("Failed to make DELETE request: " + err.Error())
	} else {
		log.Info("DELETE request successful: " + deleteBody)
		fmt.Println("DELETE Response:", deleteBody)
	}
}
