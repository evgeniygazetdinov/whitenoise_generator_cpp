package main
import (
    "io/ioutil";
    "fmt";
    "container/list";
    "strings"
    // "reflect";
    )

    func getLinksForDownload() string {
        contents,_ := ioutil.ReadFile("plikTekstowy.txt")
        return string(contents);
    }

    func makeArrayWithLinks(myString string)[]string{
        arr := strings.FieldsFunc(myString, func(r rune) bool {
           return r == ','
        })
        return arr
    }

func main() {
    // new linked list

    fmt.Println("%q\n",makeArrayWithLinks(getLinksForDownload()));
    queue := list.New()

    queue.PushBack(10)
    queue.PushBack(20)
    queue.PushBack(30)
    front:=queue.Front()
    fmt.Println(front.Value)
    queue.Remove(front)
    //TODO 1) push all links into queue
    //2)make one thread for download
    //3)zipping each into archive
    //return list with link
 
}