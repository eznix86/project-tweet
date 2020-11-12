package entity

//User is person account details
type User struct {
	ID        int // int as ID is not scalable, another stuff we called ULID/UUID
	Username  string
	Password  string
	FirstName string
	LastName  string
	Followers []User
}
