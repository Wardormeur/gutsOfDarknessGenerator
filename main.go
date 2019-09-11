package main

import (
  "log"
  "fmt"
  "strconv"
  "math/rand"
  "net/http"
  "gutsOfDarkness/configuration"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  conf := configuration.New(configuration.NewCLIParser())
  // Config generation: select note, select style; randomly
  uri := buildUrl(conf.Rank, conf.Year, 1)
  fmt.Println(uri)
  doc := loadPage(uri)
   // Find the number of pages for that selection
  var pages int = 0
  var selectedPage int = pages
  navigationContainer := doc.Find("div.barre").First()
  pages, _ = strconv.Atoi(navigationContainer.Find("strong").Last().Text())
  selectedPage = rand.Intn(pages)

  // Reload the page with all randomness included
  uri = buildUrl(conf.Rank, conf.Year, selectedPage)
  fmt.Println(uri)
  doc = loadPage(uri)
  // Select a random album
  albums := doc.Find("ul.objectList li")
  albumIndex := rand.Intn(albums.Length())
  album := albums.Slice(albumIndex, albumIndex + 1)
  band := album.Find("dt a").First().Text()
  name := album.Find("dt a").Get(1).FirstChild.Data
  fmt.Printf("%s - %s\n", band, name)
}

func buildUrl(rank int, year string, page int) string {
  return fmt.Sprintf("https://www.gutsofdarkness.com/god/selection.php?page=%v&note=%v&typeobjet=0&annee=%v&membre=&pays=0&tri=a", page, rank, year)
}

func loadPage(uri string) *goquery.Document {
  res, err := http.Get(uri)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  return doc
}
