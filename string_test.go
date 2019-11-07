package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func Test_string(t *testing.T) {
	s := "dimensions=%7Bport%3D10080%2C+protocol%3Dtcp%2C+userId%3D1854764472318598%2C+instanceId%3Dlb-t4n4zszmjw0o8to3nyedp%7D&expression=%24Maximum%3C10000"
	// var s = `alertName=%E6%B5%8B%E8%AF%95%EF%BC%88%E8%BF%9E%E6%8E%A5%E6%95%B0%E5%B0%8F%E4%BA%8E10000%EF%BC%89&alertState=ALERT&curValue=0&dimensions=%7Bport%3D10080%2C+protocol%3Dtcp%2C+userId%3D1854764472318598%2C+instanceId%3Dlb-t4n4zszmjw0o8to3nyedp%7D&expression=%24Maximum%3C10000&instanceName=tcp%EF%BC%8C%E7%AB%AF%E5%8F%A3%EF%BC%9A10080%EF%BC%8C%E5%AE%9E%E4%BE%8B%EF%BC%9Akblhx-al-tw-prod-all-game5-lb&metricName=%E7%AB%AF%E5%8F%A3%E6%B4%BB%E8%B7%83%E8%BF%9E%E6%8E%A5%E6%95%B0&metricProject=acs_slb&namespace=acs_slb&preTriggerLevel=INFO&ruleId=lb-t4n4zszmjw0o8to3nyedp_c707f7b0-cb86-403e-8181-43e4910d1399&signature=9Hv314fh1YaHT6D0wbxBcqrfuUk%3D&timestamp=157145442000s&triggerLevel=INFs&userId=1854764472318598`
	value, err := url.ParseQuery(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	res, err := http.Post("http://10.108.0.117:8080/recall", "application/x-www-form-urlencoded", bytes.NewBufferString(s))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(string(body))

}
func Test_sToJson(t *testing.T) {
	s := "{port=10081, protocol=tcp, userId=1854764472318598, instanceId=lb-t4nl80iev39mbr4bkdp79}"
	s1 := strings.Replace(s, "=", ":", 100)
	s2 := strings.Replace(s1, "{", "", 1)
	s3 := strings.Replace(s2, "}", "", 1)
	s4 := strings.Split(s3, ",")
	m := make(map[string]string)
	for _, v := range s4 {
		s5 := strings.Split(v, ":")
		m[s5[0]] = s5[1]
	}
	fmt.Println(m)
}
func Test_readIns(t *testing.T) {
	f, err := os.Open("monitor.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewReader(f)
	m := make(map[string]string)
	for {
		l, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		slice := strings.Fields(string(l))
		m[slice[1]] = slice[2]
	}
	for k, v := range m {
		fmt.Printf("metric[\"%s\"]=\"%s\"\n", k, v)
	}
}
