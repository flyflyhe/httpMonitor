package httpMonitor

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func Monitor(monitorURL string) (result map[string]string, err error) {
	result = make(map[string]string)
	proxyArr, err := GetByBucket(BucketProxy)

	if err != nil {
		return
	}
	proxyArr = append([]string{""}, proxyArr...)
	wg := sync.WaitGroup{}
	wg.Add(len(proxyArr))

	resultChan := make(chan [2]string, len(proxyArr))
	for _, proxy := range proxyArr {
		go func(proxy string) {
			defer wg.Done()
			err := send(monitorURL, proxy)
			msg := "success"
			if err != nil {
				msg = err.Error()
			}
			resultChan <- [2]string{proxy, msg}
		}(proxy)
	}

	wg.Wait()
	close(resultChan)
	for v := range resultChan {
		result[v[0]] = v[1]
	}

	return
}

func send(monitorURL string, proxy string) error {
	request, err := http.NewRequest(http.MethodGet, monitorURL, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	transport := &http.Transport{}

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return err
		}

		transport.Proxy = http.ProxyURL(proxyURL)
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}

	if res.StatusCode > 500 {
		return errors.New(res.Status)
	}

	return nil
}
