package postex
//Contains helper functions to wrap the functionality in net/http for sending and receiving data over HTTPS.

import "net/http"
import "crypto/tls"
import "bytes"
import "io/ioutil"

func getClient(skip_verify bool) *http.Client {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: skip_verify}}
	client := &http.Client{Transport: tr}
	return client
}

func doPost(url string, postdata []byte, client *http.Client) error {
	req,err := http.NewRequest("POST", url, bytes.NewBuffer(postdata))
	if err != nil {
		return err
	}
	resp,err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func doGet(url string, client *http.Client) (string,error) {
	req,err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "",err
	}
	resp,err := client.Do(req)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body),nil	
}
