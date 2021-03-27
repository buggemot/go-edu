package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"net"
	"github.com/buggemot/go-edu/pkg/csvfile"
)


func getPage(url string) (int, error) {

	t := &http.Transport{
            Dial: (&net.Dialer{
                    Timeout: 60 * time.Second,
                    KeepAlive: 30 * time.Second,
            }).Dial,
            // We use ABSURDLY large keys, and should probably not.
            TLSHandshakeTimeout: 60 * time.Second,
    }
    c := &http.Client{
            Transport: t,
    }


	resp, err := c.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}
 
func emit(c chan string, done chan bool) {
	words := []string{"This", "is", "test", "message", "hope", "true"}
	i := 0
	for {
		select {
		case c <-words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done:
			close(done)
			return 
		}
	}
}


func worker(urlCh chan string, sizeCh chan string, id int) {

	for {
		url := <- urlCh
		length, err := getPage(url)

		if err == nil {
			sizeCh <- fmt.Sprintf("%s has length %d (%d)", url, length, id)
		} else {
			sizeCh <- fmt.Sprintf("err %s: %s", url, err)
		}
	}
}


func generator(url string, urlCh chan string) {
	urlCh <- url
}


func main() {
	urls := []string{"http://google.ru", "http://habr.com/",
					 "http://yandex.ru", "http://yahoo.com",
					 "http://mail.ru", "http://bus.gov.ru",
					 "http://roskazna.ru", "http://gov.ru",
					}
	
	urlCh := make(chan string)
	sizeCh := make(chan string)

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	for _, url := range urls{
		go generator(url, urlCh)	
	}
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

	fmt.Printf("%v", csvfile.Read)
	//cf := csvfile.NewCsvFile() 
	//cf.name = "pattern.csv"
	//cf.Read()
	//fmt.Printf("%s\n %v", cf.name, cf.records[1][0])
}