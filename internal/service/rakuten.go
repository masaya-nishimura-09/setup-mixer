package service

import (
	"github.com/masaya-nishimura-09/setup-mixer/internal/model"
	"github.com/masaya-nishimura-09/setup-mixer/internal/repository"
)

func SearchItems(q string)  *model.Items {
    var items model.Items
    items.Desk = repository.Search(215698, q + " オフィスデスク")
    items.Chair = repository.Search(111363, q + " オフィスチェア")
    items.Keyboard = repository.Search(560088, q + " キーボード")
    items.Mouse = repository.Search(565170, q + " マウス")
    items.MousePad = repository.Search(552391, q + " マウスパッド")
    items.Monitor = repository.Search(110105, q + " モニター")
    items.MonitorArm = repository.Search(566221, q + " モニターアーム")
    items.DeskLamp = repository.Search(500281, q + " デスクランプ")
    items.PowerStrip = repository.Search(552481, q + " 電源タップ")
    items.PowerStripCase = repository.Search(200166, q + " ケーブルボックス")
    items.Speaker = repository.Search(208316, q + " スピーカー")
    items.Camera = repository.Search(403512, q + " ウェブカメラ")
    items.Headphone = repository.Search(568359, q + " ヘッドセット")

    return  &items
}


