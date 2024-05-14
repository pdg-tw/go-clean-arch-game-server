package member

// Repository Interface for members
type Repository interface {
	Add(mem Member) error
}
