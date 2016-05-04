package main

import(
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "io/ioutil"
  "os"
  "encoding/json"
  "time"
)
type Mesi struct{
   Day time.Time `json:"id"`
   Launch string `json:"text"`
   Don string `json:"don"`
   Salad string `json:"salad"`


}
func main(){
  i :=0
  var d [5]Mesi
  doc, _ := goquery.NewDocument("http://www.gakushoku.com/juniorcollege/daily_special.php?entry_id=2016-05-02")

  doc.Find("td").Each(func(_ int, s *goquery.Selection) {
    url, _ := s.Attr("href")
    if url != "#listtop" {
      keyword := s.Text()
        if len(keyword) > 0 {

          d[0].Day = time.Date(2016,5,2,0,0,0,0,time.Local)
          d[1].Day = d[0].Day.AddDate(0,0,1)
          d[2].Day = d[0].Day.AddDate(0,0,2)
          d[3].Day = d[0].Day.AddDate(0,0,3)
          d[4].Day = d[0].Day.AddDate(0,0,4)
            if i < 5{

            d[i].Launch = keyword

          }
            if i >= 5 && 10>i {
              d[i-5].Don = keyword

            }
            if i >= 10 && i < 15{
              d[i-10].Salad = keyword
            }


            fmt.Println(i)
            i++
          //  fmt.Println(d)
            bytes, _ := json.Marshal(d)
         fmt.Println(keyword)

           ioutil.WriteFile("./tandai.json", bytes, os.ModePerm)
        }
    }
})
}
