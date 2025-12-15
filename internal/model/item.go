package model

type ItemPage struct{
    Nav []ItemNav
    Categories []Category
}

type ItemNav struct {
    Link string
    Name string
}

type Category struct {
    NameJa string
    NameEn string
    Items []Item
}

type SearchResult struct {
    GenreInformation []int  `json:"GenreInformation"`
    Items            []Item `json:"Items"`
    TagInformation   []int  `json:"TagInformation"`
    Carrier          int    `json:"carrier"`
    Count            int    `json:"count"`
    First            int    `json:"first"`
    Hits             int    `json:"hits"`
    Last             int    `json:"last"`
    Page             int    `json:"page"`
    PageCount        int    `json:"pageCount"`
}

type Item struct {
    Availability    int      `json:"availability"`
    Catchcopy       string   `json:"catchcopy"`
    ItemCaption     string   `json:"itemCaption"`
    ItemCode        string   `json:"itemCode"`
    ItemName        string   `json:"itemName"`
    ItemPrice       int      `json:"itemPrice"`
    ItemUrl         string   `json:"itemUrl"`
    MediumImageUrls []string `json:"mediumImageUrls"`
    ReviewAverage   float64  `json:"reviewAverage"`
    ReviewCount     int      `json:"reviewCount"`
    ShopName        string   `json:"shopName"`
    ShopCode        string   `json:"shopCode"`
}
