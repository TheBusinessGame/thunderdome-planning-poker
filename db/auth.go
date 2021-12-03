package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/StevenWeathers/thunderdome-planning-poker/model"
)

// AuthUser authenticate the user
func (d *Database) AuthUser(UserEmail string, UserPassword string) (*model.User, error) {
	var user model.User
	var passHash string
	var UserLocale sql.NullString

	e := d.db.QueryRow(
		`SELECT id, name, email, type, password, avatar, verified, notifications_enabled, locale FROM users WHERE email = $1`,
		UserEmail,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Type,
		&passHash,
		&user.Avatar,
		&user.Verified,
		&user.NotificationsEnabled,
		&UserLocale,
	)
	if e != nil {
		log.Println(e)
		return nil, errors.New("user not found")
	}

	if !comparePasswords(passHash, UserPassword) {
		return nil, errors.New("password invalid")
	}

	// check to see if the bcrypt cost has been updated, if not do so
	if checkPasswordCost(passHash) == true {
		hashedPassword, hashErr := hashSaltPassword(UserPassword)
		if hashErr == nil {
			d.db.Exec(`call update_user_password($1, $2)`, user.Id, hashedPassword)
		}
	}

	user.Locale = UserLocale.String

	return &user, nil
}

// UserResetRequest inserts a new user reset request
func (d *Database) UserResetRequest(UserEmail string) (resetID string, UserName string, resetErr error) {
	var ResetID sql.NullString
	var UserID sql.NullString
	var name sql.NullString

	e := d.db.QueryRow(`
		SELECT resetId, userId, userName FROM insert_user_reset($1);
		`,
		UserEmail,
	).Scan(&ResetID, &UserID, &name)
	if e != nil {
		log.Println("Unable to reset user: ", e)
		return "", "", e
	}

	return ResetID.String, name.String, nil
}

// UserResetPassword resets the user's password to a new password
func (d *Database) UserResetPassword(ResetID string, UserPassword string) (UserName string, UserEmail string, resetErr error) {
	var name sql.NullString
	var email sql.NullString

	hashedPassword, hashErr := hashSaltPassword(UserPassword)
	if hashErr != nil {
		return "", "", hashErr
	}

	UserErr := d.db.QueryRow(`
		SELECT
			w.name, w.email
		FROM user_reset wr
		LEFT JOIN users w ON w.id = wr.user_id
		WHERE wr.reset_id = $1;
		`,
		ResetID,
	).Scan(&name, &email)
	if UserErr != nil {
		log.Println("Unable to get user for password reset confirmation email: ", UserErr)
		return "", "", UserErr
	}

	if _, err := d.db.Exec(
		`call reset_user_password($1, $2)`, ResetID, hashedPassword); err != nil {
		return "", "", err
	}

	return name.String, email.String, nil
}

// UserUpdatePassword updates a users password
func (d *Database) UserUpdatePassword(UserID string, UserPassword string) (Name string, Email string, resetErr error) {
	var UserName sql.NullString
	var UserEmail sql.NullString

	UserErr := d.db.QueryRow(`
		SELECT
			w.name, w.email
		FROM users w
		WHERE w.id = $1;
		`,
		UserID,
	).Scan(&UserName, &UserEmail)
	if UserErr != nil {
		log.Println("Unable to get user for password update: ", UserErr)
		return "", "", UserErr
	}

	hashedPassword, hashErr := hashSaltPassword(UserPassword)
	if hashErr != nil {
		return "", "", hashErr
	}

	if _, err := d.db.Exec(
		`call update_user_password($1, $2)`, UserID, hashedPassword); err != nil {
		return "", "", err
	}

	return UserName.String, UserEmail.String, nil
}

// VerifyUserAccount updates a user account verified status
func (d *Database) VerifyUserAccount(VerifyID string) error {
	if _, err := d.db.Exec(
		`call verify_user_account($1)`, VerifyID); err != nil {
		return err
	}

	return nil
}
