package tools

import (
	"net/url"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"net/http/cookiejar"
	"compress/gzip"
	"regexp"
)

const YT_URL = "http://www.youtube.com/get_video_info?hl=en_US&el=detailpage&video_id="

type Info struct {
    Title string
    Author string
    UseSign bool
    Urls []string
}

func getVideoId(urlString string) string {
    u, err := url.Parse(urlString)
    if err != nil {
       log.Fatal(err)
    }
    values := u.Host()
    return values
}

func GetInfo(urlString string) (youtube Info, err error) {
	urlReq := fmt.Sprintf("%s%s", YT_URL, getVideoId(urlString))
	if resp, err := http.Get(urlReq);err != nil {
           log.Fatal(err)
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
	    log.Fatal(err)
	}

	if link, err := url.ParseQuery(string(body));err != nil{
	    log.Fatal(err)
	}

	//fmt.Println(info)
	youtube.Title(link.Get("title"))
	youtube.Author(link.Get("author"))

	if link.Get("use_cipher_signature") == "True" {
		youtube.UseSign = true
	} else {
		youtube.UseSign = false
	}
	if youtube.UseSign {
		videoUrl := getVideoUrlFromAnotherSite(urlString)
		youtube.Urls = append(youtube.Urls, videoUrl)
		return youtube
	}

	streams_raw := strings.Split(link.Get("url_encoded_fmt_stream_map"), ",")
	for _, streams := range streams_raw {
		u, err = url.ParseQuery(streams)
		if err != nil {
			log.Fatal(err)
		}
		for _, url := range u["url"] {
			youtube.Urls = append(youtube.Urls, url)
		}
	}
	return
}


func getVideoUrlFromAnotherSite(urlString string) string {
	downloaderUrl := "http://9xbuddy.com/youtube?url="
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client {
		Jar: cookieJar,
	}

	resp, err := client.Get(fmt.Sprintf("%s%s", downloaderUrl, urlString))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	req, _ := http.NewRequest("POST",
		"http://9xbuddy.com/includes/main-post.php",
		strings.NewReader(fmt.Sprintf("url=%s", url.QueryEscape(urlString))))
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8,zh-TW;q=0.6,zh;q=0.4,zh-CN;q=0.2")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Host", "9xbuddy.com")
	req.Header.Add("Origin", "http://9xbuddy.com")
	req.Header.Add("Referer", fmt.Sprintf("%s%s", downloaderUrl, urlString))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36")


	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader, _ := gzip.NewReader(resp.Body)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	string_body := string(body)
	re := regexp.MustCompile("\"downtube.php.+?\"")
	urls := re.FindAllString(string_body, -1)
	return "http://9xbuddy.com/" + strings.Trim(urls[1], "\"")
}

