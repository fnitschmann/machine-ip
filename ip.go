package ip

import (
	"errors"
	"github.com/jpillora/backoff"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

// Get the IP for the machine in the local network
// ATTENTION: Just uses the first network interface
func GetLocalMachineIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Couldn't determine local machine IP")
}

// Get the public machine IP via ipify service (http://ipify.org)
// It uses an exponential backoff by default and if it fails it will be retried
// 3 times
func GetPublicMachineIp() (string, error) {
	b := backoff.Backoff{
		Jitter: true,
	}
	client := http.Client{}

	req, err := http.NewRequest("GET", API_URI, nil)
	if err != nil {
		return "", errors.New("Received an invalid status code from ipify.org: 500")
	}

	for tries := 0; tries < MAX_TRIES; tries++ {
		resp, err := client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", errors.New("Received an invalid status code from ipify.org: 500")
		}

		if resp.StatusCode != 200 {
			return "", errors.New("Received an invalid status code from ipify.org: " + strconv.Itoa(resp.StatusCode))
		}

		return string(ip), nil
	}

	return "", errors.New("Request failed because it wasn't able to reach the ipify service")
}
