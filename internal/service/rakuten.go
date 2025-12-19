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
		Items:  repository.Search(215698, "オフィスデスク "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "chair",
		NameJa: "イス",
		Items:  repository.Search(111363, "オフィスチェア "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "keyboard",
		NameJa: "キーボード",
		Items:  repository.Search(560088, "キーボード "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "mouse",
		NameJa: "マウス",
		Items:  repository.Search(565170, "マウス "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "mouse-pad",
		NameJa: "マウスパッド",
		Items:  repository.Search(0, "マウスパッド "+q),
	})
	//   categories = append(categories, model.Category{
	//       NameEn: "mouse-pad",
	//       NameJa: "マウスパッド",
	//       Items: repository.Search(552391, q + " マウスパッド"),
	//   })
	categories = append(categories, model.Category{
		NameEn: "monitor",
		NameJa: "モニター",
		Items:  repository.Search(110105, "モニター "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "monitor-arm",
		NameJa: "モニターアーム",
		Items:  repository.Search(566221, "モニターアーム "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "desk-lamp",
		NameJa: "デスクランプ",
		Items:  repository.Search(500281, "デスクランプ オフィス "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "power-strip",
		NameJa: "電源タップ",
		Items:  repository.Search(552481, "電源タップ 延長コード "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "power-strip-case",
		NameJa: "ケーブル収納",
		Items:  repository.Search(200166, "ケーブルボックス "+q),
	})
	//categories = append(categories, model.Category{
	//    NameEn: "speaker",
	//    NameJa: "スピーカー",
	//    Items: repository.Search(208316, q + " pcスピーカー"),
	//})
	categories = append(categories, model.Category{
		NameEn: "speaker",
		NameJa: "スピーカー",
		Items:  repository.Search(208316, "スピーカー "+q),
	})
	categories = append(categories, model.Category{
		NameEn: "camera",
		NameJa: "ウェブカメラ",
		Items:  repository.Search(403512, "ウェブカメラ "+q),
	})
	//   categories = append(categories, model.Category{
	//       NameEn: "camera",
	//       NameJa: "ウェブカメラ",
	//       Items: repository.Search(403512, q + " ウェブカメラ"),
	//   })
	categories = append(categories, model.Category{
		NameEn: "headphone",
		NameJa: "ヘッドセット",
		Items:  repository.Search(0, "ヘッドセット "+q),
	})
	//    categories = append(categories, model.Category{
	//        NameEn: "headphone",
	//        NameJa: "ヘッドセット",
	//        Items: repository.Search(568359, q + " ヘッドセット"),
	//    })
	return categories
}
