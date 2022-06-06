package checker

import (
	"fmt"
	"net"
	"time"
)

func GetChecker(timeout int) func(string) (bool, error) {
	return func(ip string) (bool, error) {
		return checker(ip, time.Duration(timeout)*time.Millisecond)
	}
}

func checker(ip string, timeout time.Duration) (bool, error) {
	/*
	  Put here some logic, that identifies your server and has very low chance
	  of matching someone else's. It might be checking for some unique pattern
	  of opened ports or some substring in html content of website or even attempt
	  of ssh connection. You are not limited by anything
	*/

	return false, nil
}

func isPortOpened(address string, port int, timeout time.Duration) bool {
	address_port := fmt.Sprintf("%s:%d", address, port)
	conn, err := net.DialTimeout("tcp", address_port, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
