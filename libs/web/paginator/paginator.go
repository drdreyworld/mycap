package paginator

import "fmt"

type Page struct {
	Url     string
	Title   string
	Current bool
}

type Pages []Page

type Items interface {
	Len() int
}

type Paginator struct {
	Page  int
	Count int
	Scale int
	Pages Pages

	ItemsStart int
	ItemsStop  int
	IsNeed     bool
}

func (self *Paginator) Create(page int, scale int, items Items) {
	self.Scale = scale
	self.Count = items.Len() / self.Scale
	if items.Len()%self.Scale > 0 {
		self.Count++
	}

	self.Page = page

	if self.Page < 1 || self.Page > self.Count {
		self.Page = 1
	}

	if self.IsNeed = self.Count > 1; self.IsNeed {
		self.Pages = Pages{}

		num := 0
		max := 3

		if self.Page < 6 && page > 2 {
			max = self.Page + 1
		}

		for i := 1; i <= self.Count && i <= max; i++ {
			num = i
			self.Pages = append(self.Pages, self.CreatePage(i))
		}

		if self.Page >= 6 && self.Page < self.Count-1 {
			self.Pages = append(self.Pages, self.CreatePageSeparator())
			max = self.Page + 1
			for i := self.Page - 1; i <= self.Count-1 && i <= max; i++ {
				num = i
				self.Pages = append(self.Pages, self.CreatePage(i))
			}
		}

		num++

		if num < self.Count-2 {
			self.Pages = append(self.Pages, self.CreatePageSeparator())
			num = self.Count - 2
		}

		for i := num; i <= self.Count; i++ {
			num = i
			self.Pages = append(self.Pages, self.CreatePage(i))
		}
	}

	self.ItemsStart = self.Scale * (self.Page - 1)
	self.ItemsStop = self.Scale * self.Page

	if self.ItemsStop > items.Len() {
		self.ItemsStop = items.Len()
	}
}

func (self *Paginator) CreatePage(i int) Page {
	page := Page{
		Url:     fmt.Sprintf("?page=%d", i),
		Title:   fmt.Sprintf("%d", i),
		Current: i == self.Page,
	}
	return page
}

func (self *Paginator) CreatePageSeparator() Page {
	page := Page{
		Url:     "",
		Title:   "...",
		Current: false,
	}
	return page
}
