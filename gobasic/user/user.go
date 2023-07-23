package user

type Person struct{
	name string
	age int
}


func (p Person) GetName() string{
	return p.name
}

func (p *Person) SetName(name string){
	p.name = name
}

func (p Person) GetAge()int{
	return p.age
}

func (p Person) Hello() string{
	return "Hello" + p.name
}