package repository

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "time"

	"github.com/masaya-nishimura-09/setup-mixer/internal/model"
)

func Search(genreId int, keywords string) []model.Item {
    time.Sleep(200 * time.Millisecond)

    baseURL := "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20220601"
    u, err := url.Parse(baseURL)
    if err != nil {
        fmt.Println(err)
        return []model.Item{}
    }
    values := u.Query()
    values.Add("keyword", keywords)
    values.Add("hasReviewFlag", "1")
    values.Add("format", "json")
    values.Add("sort", "-reviewCount")
    values.Add("carrier", "0")
    values.Add("availability", "1")
    values.Add("field", "1")
    values.Add("imageFlag", "1")
    values.Add("formatVersion", "2")
    values.Add("applicationId", os.Getenv("APPLICATION_ID"))
    values.Set("genreId", strconv.Itoa(genreId))
    u.RawQuery = values.Encode()

    req, err := http.NewRequest(http.MethodGet, u.String(), nil)
    if err != nil {
        fmt.Println(err)
    }

    client := new(http.Client)
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }

    var searchResult model.SearchResult
    if err := json.Unmarshal(body, &searchResult); err != nil {
        fmt.Println(err)
    }

    fmt.Println(keywords)

    return searchResult.Items
}
