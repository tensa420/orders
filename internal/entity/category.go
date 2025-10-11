package entity

type Category string

const (
	CategoryUnspecified Category = "UNSPECIFIED" // Категория не указана.
	CategoryEngine      Category = "ENGINE"      // Двигатели и компоненты.
	CategoryFuel        Category = "FUEL"        // Топливная система.
	CategoryPorthole    Category = "PORTHOLE"    // Иллюминаторы и окна.
	CategoryWing        Category = "WING"        // Крылья и аэродинамические поверхности.
)
