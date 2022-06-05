package httpMonitor

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Monitor(monitorURL string) (result map[string]string, err error) {
	result = make(map[string]string)
	proxyArr, err := GetUrlProxyList(monitorURL)

	if err != nil {
		return
	}

	proxyArr = append([]string{""}, proxyArr...)

	for _, proxy := range proxyArr {
		err = send(monitorURL, proxy)
		if err != nil {
			result[proxy] = err.Error()
		} else {
			result[proxy] = "success"
		}
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
