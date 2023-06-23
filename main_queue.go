package main
import (
    "io/ioutil";
    "fmt";
    "container/list";
    "strings";
    "math/rand"
    "log";
    "net/http";
    "reflect"
    )

    func getLinksForDownload() string {
        contents,_ := ioutil.ReadFile("examples.txt")
        return string(contents);
    }

    func makeArrayWithLinks(myString string)[]string{
        arr := strings.FieldsFunc(myString, func(r rune) bool {
           return r == ','
        })
        return arr
    }
    func prepareUrls(links []string)*list.List{
        queue := list.New()
        for _, link := range links {
            queue.PushBack(link)
          }
        return queue
    }
    func getDataForQueue()*list.List{
       arrayWithUrl := makeArrayWithLinks(getLinksForDownload())
       queue := prepareUrls(arrayWithUrl) 
       return queue 
    }

    func awaitTask(myValue string, myUrl string) <-chan string {
        fmt.Println(myValue)
        fmt.Println("Starting Task...")
    
        c := make(chan string)
    
        go func() {
            resp, err := http.Get(myUrl)
            if err != nil {
                log.Fatalln(err)
            }
    
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                log.Fatalln(err)
            }
    
            c <- string(body)
    
            fmt.Println("...Done!")
        }()
    
        return c
    }

func getMethods(myQueue *list.List){
    fooType := reflect.TypeOf(myQueue)
    for i := 0; i < fooType.NumMethod(); i++ {
        method := fooType.Method(i)
        fmt.Println(method.Name)

    }
}
    
func RandStringBytes(n int) string {
    const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

// func downloadMultipleFiles(urls []string) ([][]byte, error) {
// 	done := make(chan []byte, len(urls))
// 	errch := make(chan error, len(urls))
// 	for _, URL := range urls {
// 		go func(URL string) {
// 			b, err := downloadFile(URL)
// 			if err != nil {
// 				errch <- err
// 				done <- nil
// 				return
// 			}
// 			done <- b
// 			errch <- nil
// 		}(URL)
// 	}
// 	bytesArray := make([][]byte, 0)
// 	var errStr string
// 	for i := 0; i < len(urls); i++ {
// 		bytesArray = append(bytesArray, <-done)
// 		if err := <-errch; err != nil {
// 			errStr = errStr + " " + err.Error()
// 		}
// 	}
// 	var err error
// 	if errStr!=""{
// 		err = errors.New(errStr)
// 	}
// 	return bytesArray, err
// }

func main() {
    myQueue := getDataForQueue()

    getMethods(myQueue)
    for i:=0;i<myQueue.Len();i++{
        myRandomString := RandStringBytes(i)
            value := <-awaitTask(myRandomString, myQueue.Front())
            myQueue.Remove()
            fmt.Println(value)
          }
    }
