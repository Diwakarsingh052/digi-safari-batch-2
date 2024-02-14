package data

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(data string)
	Update(data string)
	Delete(data string)
}

// Store manages the set of APIs for data access.
type Store struct {
	Storer
}

// NewStore constructs a Storer for data api access.
func NewStore(s Storer) Store {
	return Store{Storer: s}
}
