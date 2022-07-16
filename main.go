package main

type Getter interface {
	Get() string
}

type Setter interface {
	Set(string)
}

type AbstractGetter struct {
	Name string
}

func (g *AbstractGetter) Set(Name string) {
	g.Name = Name
}

type NoticeGetter struct {
	AbstractGetter
}

func (g *NoticeGetter) Get() string {
	return "NoticeGetter"
}

func main() {

}
