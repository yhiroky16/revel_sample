package models

type ViewSearchResult struct {
	Number string
	Name string
	Affilation string
	Responsible string
	Mail string
	Tel string
	Entry_month string
	Birthday string
}

func (view *ViewSearchResult) SetView(number string, name string, affilation string, responsible string,
	mail string, tel string, entry_month string, birthday string) {

	view.Number = number
	view.Name = name
	view.Affilation = affilation
	view.Responsible = responsible
	view.Mail = mail
	view.Tel = tel
	view.Entry_month = entry_month
	view.Birthday = birthday
}