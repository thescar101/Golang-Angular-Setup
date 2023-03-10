package newsfeed

import "gorm.io/gorm"

// type Getter interface {
// 	GetAll() []Item
// }

// type Adder interface {
// 	Add(item Item)
// }

type Item struct {
	gorm.Model
	Title string `json:"title"`
	Post  string `json:"post"`
}

// type Repo struct {
// 	Items []Item
// }

// func New() *Repo {
// 	return &Repo{
// 		Items: []Item{
// 			{
// 				"Hello",
// 				"World",
// 			},
// 		},
// 	}
// }

// func (r *Repo) Add(item Item) {
// 	r.Items = append(r.Items, item)
// }

// func (r *Repo) GetAll() []Item {
// 	return r.Items
// }
