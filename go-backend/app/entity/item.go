package entity

import (
	"fmt"
	"fs/lib"
	"log"
)

type Item struct {
	Id       string `json:"id" gorm:"id,omitempty"`
	ItemType int    `json:"type" gorm:"item_type"`
	OrderNum int    `json:"order" gorm:"order_num"`
	ParentId string `json:"parent_id" gorm:"parent_id,omitempty"`
	Name     string `json:"name" gorm:"name"`
	Data     string `json:"data" gorm:"data,omitempty"`
}

func (g *Item) CheckValid() error {
	if lib.CheckName(g.Name) == false {
		log.Printf("error at checking name: %s\n", g.Name)
		return ErrorNameNotValid
	} else if g.ItemType == 0 && len(g.Data) != 0 {
		log.Printf("error folder tiem with data - %v\n", g)
		return ErrorFolderCantContainData
	}
	return nil
}

func (g *Item) TableName() string {
	return "items"
}

func (g *Item) ToString() string {
	return fmt.Sprintf("Id: %s, Name: %s, ParentId: %s, Order: %d\n", g.Id, g.Name, g.ParentId, g.OrderNum)
}
