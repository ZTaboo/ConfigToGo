package model

type ResData struct {
	Name  string
	Param *[]Param
}

type Param struct {
	Name    string
	Type    string
	Tag     string
	TagType string
}
