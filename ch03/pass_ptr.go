package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User  //传递指针，使用前必须先初始化才能用
	GameId int
}

func main() {
	p := &Player{User: &User{42, "Matt", "LA"}, GameId: 1}
	// 等同于下面初始化Player
	//p := &Player{&User{42, "Matt", "LA"}, 1}
	p.Name = "Mike"
	p.Location = "China"
	p.User.Id = 11
	fmt.Println(p.Greetings())
}
