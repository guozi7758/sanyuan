package lib

import (
	"bytes"
	"strings"
)
import "net/http"
import "io/ioutil"

func XFormPost( str_url, para string )( []byte, error ){
	client := &http.Client{}
	req,err := http.NewRequest("POST",str_url,strings.NewReader( para ))
	if err != nil {
		return make([]byte, 0), err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err_a := client.Do(req)
	if err_a != nil {
		return make([]byte, 0), err_a
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func PayFormPost( str_url, para string )( []byte, error ){
	client := &http.Client{}
	req,err := http.NewRequest("POST",str_url,strings.NewReader( para ))
	if err != nil {
		return make([]byte, 0), err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err_a := client.Do(req)
	if err_a != nil {
		return make([]byte, 0), err_a
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func PayJsonPost( str_url, para string )( []byte, error ){
	client := &http.Client{}
	var jsonStr = []byte(para)
	req,err := http.NewRequest("POST",str_url,bytes.NewBuffer(jsonStr))
	if err != nil {
		return make([]byte, 0), err
	}
	req.Header.Set("Content-Type",  "application/json")
	resp, err_a := client.Do(req)
	if err_a != nil {
		return make([]byte, 0), err_a
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

