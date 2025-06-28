package pkg

type Item struct {
	ItemTitle     string `json:"name"`
	ItemCompleted bool   `json:"completed"`
}

func NewItem(title string) Item    { return Item{ItemTitle: title} }
func (i Item) Completed() bool     { return i.ItemCompleted }
func (i Item) Title() string       { return i.ItemTitle }
func (i Item) FilterValue() string { return i.ItemTitle }
