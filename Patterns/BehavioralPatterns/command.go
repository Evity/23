package main

type Doctor struct {
}

func (d *Doctor) treatEyes() {
	println("Doctor treat eyes")
}

func (d *Doctor) treatNose() {
	println("Doctor treat nose")
}

type CommandTreatEyes struct {
	doctor *Doctor
}

func (c *CommandTreatEyes) Treat() {
	c.doctor.treatEyes()
}

type CommandTreatNose struct {
	doctor *Doctor
}

func (c *CommandTreatNose) Treat() {
	c.doctor.treatNose()
}

func main() {
	doctor := new(Doctor)
	eyes := CommandTreatEyes{doctor}
	nose := CommandTreatNose{doctor}
	eyes.Treat()
	nose.Treat()
}
