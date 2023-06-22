package main
import (
    "io/ioutil";
    "fmt";
    "container/list";
    "strings";
    // "reflect";
    // "sync";
    "math/rand"
    "log";
    "net/http";
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
        // fmt.Println(reflect.TypeOf(queue))
        return queue
    }
    func getDataForQueue()*list.List{
       arrayWithUrl := makeArrayWithLinks(getLinksForDownload())
       queue := prepareUrls(arrayWithUrl) 
       return queue 
    }

    func awaitTask(myValue string) <-chan string {
        fmt.Println(myValue)
        fmt.Println("Starting Task...")
    
        c := make(chan string)
    
        go func() {
            resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
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
    const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func main() {
    myQueue := getDataForQueue()
    fmt.Println("%q",myQueue)

    // TODO 
    //1)make one thread for download
    //3)zipping each into archive
    //return list with link
    
    // put leng of que here
    for i:=0;i<10;i++{
        myRandomString := RandStringBytes(i)
        value := <-awaitTask(myRandomString)

        fmt.Println(value)
    }

}