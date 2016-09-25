package repositories

import "github.com/deputadosemfoco/users/domain"

// SQLUserRepository is a MySQL implementation for UserRepository
type SQLUserRepository struct{}

// Save the User to data store
func (repo *SQLUserRepository) Save(user domain.User) {
	db := Connect()
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
func (repo *SQLUserRepository) FindByEmail(email string) domain.OptionalUser {
	db := Connect()
	sql := "SELECT id, name, email, photoUrl, facebookId, googleId, createdAt, updatedAt FROM user WHERE email = ?"
	user := domain.User{}

	err := db.Get(user, sql, email)

	if err != nil {
		return domain.OptionalUser{Valid: false}
	}

	return domain.OptionalUser{Valid: true, User: user}
}
