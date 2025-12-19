package service

import (
    "github.com/masaya-nishimura-09/setup-mixer/internal/model"
    "github.com/masaya-nishimura-09/setup-mixer/internal/repository"
)

func SearchItems(q string) []model.Category {
    var categories []model.Category 

    categories = append(categories, model.Category{
        NameEn: "desk", 
        NameJa: "デスク", 
        Items: repository.Search(215698, q + " オフィスデスク"),
    })
    categories = append(categories, model.Category{
        NameEn: "chair", 
        NameJa: "イス", 
        Items: repository.Search(111363, q + " オフィスチェア"),
    })
    categories = append(categories, model.Category{
        NameEn: "keyboard", 
        NameJa: "キーボード", 
        Items: repository.Search(560088, q + " キーボード"),
    })
    categories = append(categories, model.Category{
        NameEn: "mouse",
        NameJa: "マウス", 
        Items: repository.Search(565170, q + " マウス"),
    })
    return  categories
}
