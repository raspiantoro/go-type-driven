package roles

type Admin struct{}
type Manager struct{}
type Staff struct{}

type Roles interface {
	Admin | Manager | Staff
}
