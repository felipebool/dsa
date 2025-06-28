package element

type Getter interface {
	GetKey() int
}

type Setter interface {
	SetKey(key int)
}

type GetterSetter interface {
	Getter
	Setter
}
