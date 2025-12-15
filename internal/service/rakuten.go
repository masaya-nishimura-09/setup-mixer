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
    categories = append(categories, model.Category{
        NameEn: "mouse-pad", 
        NameJa: "マウスパッド", 
        Items: repository.Search(0, q + " マウスパッド"),
    })
 //   categories = append(categories, model.Category{
 //       NameEn: "mouse-pad", 
 //       NameJa: "マウスパッド", 
 //       Items: repository.Search(552391, q + " マウスパッド"),
 //   })
    categories = append(categories, model.Category{
        NameEn: "monitor",
        NameJa: "モニター", 
        Items: repository.Search(110105, q + " モニター"),
    })
    categories = append(categories, model.Category{
        NameEn: "monitor-arm",
        NameJa: "モニターアーム", 
        Items: repository.Search(566221, q + " モニターアーム"),
    })
    categories = append(categories, model.Category{
        NameEn: "desk-lamp", 
        NameJa: "デスクランプ", 
        Items: repository.Search(500281, q + " デスクランプ オフィス"),
    })
    categories = append(categories, model.Category{
        NameEn: "power-strip",
        NameJa: "電源タップ", 
        Items: repository.Search(552481, q + " 電源タップ 延長コード"),
    })
    categories = append(categories, model.Category{
        NameEn: "power-strip-case",
        NameJa: "ケーブル収納", 
        Items: repository.Search(200166, q + " ケーブルボックス"),
    })
    categories = append(categories, model.Category{
        NameEn: "speaker", 
        NameJa: "スピーカー", 
        Items: repository.Search(208316, q + " pcスピーカー"),
    })
    categories = append(categories, model.Category{
        NameEn: "camera",
        NameJa: "ウェブカメラ", 
        Items: repository.Search(0, q + " ウェブカメラ"),
    })
 //   categories = append(categories, model.Category{
 //       NameEn: "camera",
 //       NameJa: "ウェブカメラ", 
 //       Items: repository.Search(403512, q + " ウェブカメラ"),
 //   })
    categories = append(categories, model.Category{
        NameEn: "headphone",
        NameJa: "ヘッドセット", 
        Items: repository.Search(0, q + " ヘッドセット"),
    })
//    categories = append(categories, model.Category{
//        NameEn: "headphone",
//        NameJa: "ヘッドセット", 
//        Items: repository.Search(568359, q + " ヘッドセット"),
//    })
    return  categories
}
