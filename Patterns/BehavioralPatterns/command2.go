package main

type Doctor struct {
}

func (d *Doctor) treatEyes() {
	println("Doctor treat eyes")
}

func (d *Doctor) treatNose() {
	println("Doctor treat nose")
}

type Command interface {
	Treat()
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

type Nurse struct {
	cmdList []Command
}

func (n *Nurse) Notify() {
	if n.cmdList == nil {
		return
	}
	for _, cmd := range n.cmdList {
		cmd.Treat()
	}
}

func main() {
	doctor := new(Doctor)

	eyes := CommandTreatEyes{doctor}
	nose := CommandTreatNose{doctor}

	n := new(Nurse)

	n.cmdList = append(n.cmdList, &eyes)
	n.cmdList = append(n.cmdList, &nose)

	n.Notify()
}
