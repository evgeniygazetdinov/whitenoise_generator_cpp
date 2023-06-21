package main
import (
    "io/ioutil";
    "fmt";
    "container/list";
    "strings"
    // "reflect";
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

func main() {
    myQueue := getDataForQueue()
    fmt.Println("%q",myQueue)

    // TODO 
    //1)make one thread for download
    //3)zipping each into archive
    //return list with link
 
}