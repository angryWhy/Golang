package main

type User interface {
	say(int)
	move(int)
}
type car struct {
}
type ship struct {
}

func (c car) move(a int) {

}
func (c car) say(a int) {

}
func main() {
	var t User
	t = car{}
	t.move(1)
}
