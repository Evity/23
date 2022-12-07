package main

type Ball interface {
	Bounce()
}

type BallFactory interface {
	MakeBall() Ball
}

type Basketball struct {
	//in golang you don't have to explicitly declare the inheritance
}

type Football struct {
	Ball //for understanding the inheritance
}

type Tennis struct {
	Ball //for understanding the inheritance
}

type Volleyball struct {
	Ball //for understanding the inheritance
}

type Badminton struct {
	Ball //for understanding the inheritance
}

func (b *Basketball) Bounce() {
	println("Basketball bounce")
}

func (f *Football) Bounce() {
	println("Football bounce")
}

func (t *Tennis) Bounce() {
	println("Tennis bounce")
}

func (v *Volleyball) Bounce() {
	println("Volleyball bounce")
}

func (b *Badminton) Bounce() {
	println("Badminton bounce")
}

type BasketballFactory struct {
	BallFactory
}

type FootballFactory struct {
	BallFactory
}

type TennisFactory struct {
	BF BallFactory //recommend composition instead of inheritance
}

type VolleyballFactory struct {
	BF BallFactory
}

type BadmintonFactory struct {
	BF BallFactory
}

func (bf *BasketballFactory) MakeBall() Ball {
	var ball Ball = new(Basketball)
	return ball
}

func (bf *FootballFactory) MakeBall() Ball {
	var ball Ball = new(Football)
	return ball
}

func (bf *TennisFactory) MakeBall() Ball {
	var ball Ball = new(Tennis)
	return ball
}

func (bf *VolleyballFactory) MakeBall() Ball {
	var ball Ball = new(Volleyball)
	return ball
}

func (bf *BadmintonFactory) MakeBall() Ball {
	var ball Ball = new(Badminton)
	return ball
}

func main() {
	basketballFactory := new(BasketballFactory)
	basketball := basketballFactory.MakeBall()
	basketball.Bounce()

	footballFactory := new(FootballFactory)
	football := footballFactory.MakeBall()
	football.Bounce()

	tennisFactory := new(TennisFactory)
	tennis := tennisFactory.MakeBall()
	tennis.Bounce()

	volleyballFactory := new(VolleyballFactory)
	volleyball := volleyballFactory.BF.MakeBall()
	volleyball.Bounce()

	badmintonFactory := new(BadmintonFactory)
	badminton := badmintonFactory.MakeBall()
	badminton.Bounce()
}
