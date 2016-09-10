package users

import "github.com/vitorsalgado/la-democracia/lib/go/sqldb"

type (
	// UserRepository interface for User related data access
	UserRepository interface {
		Save(user User)
		FindByEmail(email string) OptionalUser
	}

	// SQLUserRepository is a MySQL implementation for UserRepository
	SQLUserRepository struct {
	}
)

// NewUserRepository returns a new UserRepository instance
func NewUserRepository() *SQLUserRepository {
	return &SQLUserRepository{}
}

// Save the User to data store
func (repo *SQLUserRepository) Save(user User) {
	db := sqldb.Connect()
	sql := ""

	if user.ID == "" {
		sql = "INSERT INTO user (id, name, email, photoUrl, facebookId, googleId, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		db.Exec(sql, user.ID, user.Name, user.Email, user.PhotoURL, user.FacebookID, user.GoogleID, user.CreatedAt, user.UpdatedAt)
	} else {
		sql = "UPDATE user SET name = ?, email = ?, photoUrl = ?, facebookId = ?, googleId = ?, updatedAt = ? WHERE id = ?"
		db.Exec(sql, user.Name, user.Email, user.PhotoURL, user.FacebookID, user.GoogleID, user.UpdatedAt, user.ID)
	}
}

// FindByEmail searches a User by its email.
// Must receive a valid email address
// Returns a OptionalUser that may or may not contain a User inside
func (repo *SQLUserRepository) FindByEmail(email string) OptionalUser {
	db := sqldb.Connect()
	sql := "SELECT id, name, email, photoUrl, facebookId, googleId, createdAt, updatedAt FROM user WHERE email = ?"
	user := User{}

	err := db.Get(user, sql, email)

	if err != nil {
		return OptionalUser{Valid: false}
	}

	return OptionalUser{Valid: true, User: user}
}
