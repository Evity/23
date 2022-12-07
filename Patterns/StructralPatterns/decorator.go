package main

//==========abstract component==========
type Phone interface {
	Show()
}

//decorator(it should be interface, but in go, interface must not contains any field)
type PhoneDecorator struct {
	phone Phone
}

func (p *PhoneDecorator) Show() {
	p.phone.Show()
}

//==========concrete component==========

type HUAWEI struct{}

func (h *HUAWEI) Show() {
	println("HUAWEI Show")
}

type IPHONE struct{}

func (i *IPHONE) Show() {
	println("IPHONE Show")
}

//==========concrete decorator==========

type FilmDecorator struct {
	PhoneDecorator
}

func (f *FilmDecorator) Show() {
	f.PhoneDecorator.Show()
	println("FilmDecorator Show")
}

func NewFilmDecorator(phone Phone) *FilmDecorator {
	return &FilmDecorator{PhoneDecorator{phone}}
}

type ShellDecorator struct {
	PhoneDecorator
}

func (s *ShellDecorator) Show() {
	s.PhoneDecorator.Show()
	println("ShellDecorator Show")
}

func NewShellDecorator(phone Phone) *ShellDecorator {
	return &ShellDecorator{PhoneDecorator{phone}}
}

//------------logic code------------

func main() {
	var huawei Phone = new(HUAWEI)
	huawei.Show() //original method

	var filmHuawei Phone = NewFilmDecorator(huawei)
	filmHuawei.Show() //decorated method

	var shellHuawei Phone = NewShellDecorator(huawei)
	shellHuawei.Show() //decorated method

	var filmShellHuawei Phone = NewFilmDecorator(shellHuawei)
	filmShellHuawei.Show() //decorated method
}
