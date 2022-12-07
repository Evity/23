package main

type Beauty interface {
	GiveGladEye()
	Date()
}

type PanJinLian struct{}

func (p *PanJinLian) GiveGladEye() {
	println("Pan Jinlian gives you a glad eye")
}

func (p *PanJinLian) Date() {
	println("Pan Jinlian dates with you")
}

// 代理中介人, 王婆
type WangPo struct {
	beauty Beauty
}

func NewProxy(b Beauty) Beauty {
	return &WangPo{b}
}

// 对男人抛媚眼
func (p *WangPo) GiveGladEye() {
	p.beauty.GiveGladEye()
}

// 和男人浪漫的约会
func (p *WangPo) Date() {
	p.beauty.Date()
}

// 西门大官人
func main() {
	//大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinLian))
	//王婆命令潘金莲抛媚眼
	wangpo.GiveGladEye()
	//王婆命令潘金莲和西门庆约会
	wangpo.Date()
}
