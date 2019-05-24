package main

import (
    "net/http"
    "net/url"
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "strings"
)

var lengths = []string{"long", "medium", "short"}

func main() {
    paragraphs := "3"
    length := "medium"
    for i, arg := range os.Args[1:] {
        if i == 0 {
            paragraphs = arg
        } else if i == 1 {
            if contains(lengths, arg) {
                length = arg
            }
        }
    }

    // New code
    apiurl := "http://hangul.thefron.me/api/generator"
    res, err := http.PostForm(apiurl, url.Values{"utf8": {"✓"}, "paragraphs": {paragraphs}, "length": {length}, "text_source_ids[]": {"1"}, "commit": {"생성하기"}})

    if err != nil {
        fmt.Println(err)
        return
    }

    body := res.Body
    defer body.Close()

    contents, err_ := ioutil.ReadAll(body)

    if err_ != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(read(string(contents)))
}

func contains(strs []string, comp string) bool {
    for _, str := range strs { if str == comp { return true } }
    return false
}

func read(contents string) string {
    var data map[string]interface{}
    byt := []byte(contents)
    if err := json.Unmarshal(byt, &data); err != nil {
        return "Couldn't read from json!"
    }
    result := data["ipsum"].(string)
    result = strings.Replace(result, "<br><br>", "\n", -1)
    return result
}
