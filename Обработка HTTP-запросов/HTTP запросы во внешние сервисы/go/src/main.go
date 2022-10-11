package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func startServer() { // поднимаем свой небольшой  сервер, к которому будем обращаться с запросами.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getHandler: incoming request %#v\n", r)
		fmt.Fprintf(w, "getHandler: r.Url %#v\n", r.URL)
	})
	http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "postHandler: raw body %s\n", string(body))
	})
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
func RunGet() { //когда есть какойто url и нужно дёрнуть Get-запрос
	url := "http://127.0.0.1:8080/?param=123&param2=test" // адрес url с которым мы будем работать
	resp, err := http.Get(url)                            // аозрат респонс
	if err != nil {
		fmt.Println("Error happend", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)              //  дёргаем из get url
	fmt.Printf("http.Get body %#v\n\n\n", string(respBody)) // вывод
}
func runGetFullReq() { //респонс (отправка больших данных с запросом)
	req := &http.Request{ // создание структуру реквест
		Method: http.MethodGet, //метод get
		Header: http.Header{ //хедор
			"User-Agent": {"couursera/golang"},
		},
	}
	req.URL, _ = url.Parse("http://127.0.0.1:8080/?id=42") // распарсинг url и отпрвка его в респрес
	req.URL.Query().Set("user", "Roman")                   //  обращение в нулювую форму url  и потстановка значений
	resp, err := http.DefaultClient.Do(req)                // выполнения запроса и иходит в другой сервер
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)                    // вычитываем из resp.Bady
	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(respBody)) // ввывод
}
func runTransportAndPost() { // отправка пост-запроса  вместе с данными на сервер
	transport := &http.Transport{ //  создание структуры транспорт
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{ //  создание клиента
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	data := `{"id": 42, "user": "Roman"}`                     // пример данных
	body := bytes.NewBufferString(data)                       // создание io reader
	url := "http://127.0.0.1:8080/raw_body"                   // адрес url
	req, _ := http.NewRequest(http.MethodPost, url, body)     // возращения запроса (сздание новый репрест с праметрами )
	req.Header.Add("Content-Type", "application/json")        // контент type
	req.Header.Add("Content-Length", strconv.Itoa(len(data))) //контент len (сколько нужно данных передать)
	resp, err := client.Do(req)                               // вызов метод do  и передача репрес
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)             // вычитываем bady
	fmt.Printf("runTransport %#v\n\n\n", string(respBody)) // вывод
}

func main() {
	go startServer()
	time.Sleep(100 * time.Millisecond)
	RunGet()
	runGetFullReq()
	runTransportAndPost()
}
