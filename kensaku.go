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
   Spaghetti string `json:"spaghetti"`
   Fish string `json:"fish"`
   Salad string `json:"salad"`
   Dessert string `json:"dessert"`
   One string `json:"one"`
   Noodle string `json:"noodle"`
   Supper string `json:"supper"`

}
func main(){
  i :=0
  var d [5]Mesi
  doc, _ := goquery.NewDocument("http://www.gakushoku.com/canteen/daily_special.php?entry_id=2016-05-02")

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
              d[i-10].Spaghetti = keyword

            }
            if i >= 15 &&  i < 20{
              d[i-15].Fish = keyword

            }
            if i >= 20 && i < 25 {
              d[i-20].Salad = keyword

            }
            if i >= 25 && i < 30{
              d[i-25].Dessert = keyword

            }
            if i >= 30 && i < 35{
              d[i-30].One = keyword

            }
            if i >= 35 && i < 40{
              d[i-35].Noodle = keyword

            }
            if i >= 40 && i < 45{
              d[i-40].Supper = keyword

            }
            fmt.Println(i)
            i++
          //  fmt.Println(d)
            bytes, _ := json.Marshal(d)
         fmt.Println(keyword)

           ioutil.WriteFile("./go-file.json", bytes, os.ModePerm)
        }
    }
})
}
