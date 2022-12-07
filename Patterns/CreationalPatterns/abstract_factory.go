package main

//==============================abstract layer============================
type Ball interface {
	Bounce()
}

type Ellipsoid interface {
	Kick()
}

type FeatherBall interface {
	Hit()
}

type BallFactory interface {
	CreateBall() Ball
	CreateEllipsoid() Ellipsoid
	CreateFeatherBall() FeatherBall
}

//==============================concrete layer============================
type Basketball struct{}

func (b *Basketball) Bounce() {
	println("Basketball bounce")
}

type Football struct{}

func (f *Football) Kick() {
	println("Football Kicked!")
}

type Badminton struct{}

func (b *Badminton) Hit() {
	println("Badminton Pat!")
}

type ClassicFactory struct{}

func (c *ClassicFactory) CreateBall() Ball {
	return new(Basketball)
}

func (c *ClassicFactory) CreateEllipsoid() Ellipsoid {
	return new(Football)
}

func (c *ClassicFactory) CreateFeatherBall() FeatherBall {
	return new(Badminton)
}

type ChinaBasketball struct{}

func (c *ChinaBasketball) Bounce() {
	println("China Basketball bounce")
}

type ChinaFootball struct{}

func (c *ChinaFootball) Kick() {
	println("China Football Kicked!")
}

type ChinaBadminton struct{}

func (c *ChinaBadminton) Hit() {
	println("China Badminton Pat!")
}

type ChinaFactory struct{}

func (c *ChinaFactory) CreateBall() Ball {
	return new(ChinaBasketball)
}

func (c *ChinaFactory) CreateEllipsoid() Ellipsoid {
	return new(ChinaFootball)
}

func (c *ChinaFactory) CreateFeatherBall() FeatherBall {
	return new(ChinaBadminton)
}

type ExtremeBasketball struct{}

func (e *ExtremeBasketball) Bounce() {
	println("Extreme Basketball bounce")
}

type ExtremeFootball struct{}

func (e *ExtremeFootball) Kick() {
	println("Extreme Football Kicked!")
}

type ExtremeBadminton struct{}

func (e *ExtremeBadminton) Hit() {
	println("Extreme Badminton Pat!")
}

type ExtremeFactory struct{}

func (e *ExtremeFactory) CreateBall() Ball {
	return new(ExtremeBasketball)
}

func (e *ExtremeFactory) CreateEllipsoid() Ellipsoid {
	return new(ExtremeFootball)
}

func (e *ExtremeFactory) CreateFeatherBall() FeatherBall {
	return new(ExtremeBadminton)
}

//==============================logic layer============================
func main() {
	var classicFactory = new(ClassicFactory)
	var chinaFactory = new(ChinaFactory)
	var extremeFactory = new(ExtremeFactory)

	var classicBall = classicFactory.CreateBall()
	classicBall.Bounce()

	var chinaBall = chinaFactory.CreateBall()
	chinaBall.Bounce()

	var extremeBall = extremeFactory.CreateBall()
	extremeBall.Bounce()

	var classicEllipsoid = classicFactory.CreateEllipsoid()
	classicEllipsoid.Kick()
}
