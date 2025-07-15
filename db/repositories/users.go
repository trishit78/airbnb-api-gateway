package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	//Create() error
	GetByID() (*models.User, error)
	Create() ( error)
	GetAll() ([]*models.User, error)
	DeleteByID(id int64) error

}

type UserRepositoryImpl struct {
	db *sql.DB
}





func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db:_db,
	}
}




func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	rows,err := u.db.Query("SELECT id,username,email,created_at,updated_at from users")

	if err != nil{
		return nil,err
	}

	var allUsers []*models.User

	for rows.Next(){
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil{
			fmt.Println("failed to scan user",err)
			return nil,err
		}
		allUsers = append(allUsers, user)
	}
	errr := rows.Err()
	if errr != nil{
		fmt.Println("got row error",errr)
		return nil,errr
	}
	//print allUsers
	for _, user := range allUsers {
		fmt.Printf("ID: %d | Username: %s | Email: %s | CreatedAt: %s | UpdatedAt: %s\n",
			user.Id, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)
	}
	return  allUsers,nil
	
}

func (u *UserRepositoryImpl) DeleteByID(id int64) error {
	query := "delete from users where id= ?"

	result, err := u.db.Exec(query, 1)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}
	fmt.Println("User deleted successfully, rows affected:", rowsAffected)
	return nil
}




func (u *UserRepositoryImpl) Create() (error) {
query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, "testuser", "test@test.com", "password123")

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected:", rowsAffected)

	return nil
}


func (u *UserRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Getching user in UserRepository")

	// Step 1: Prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, 1)

	// Step 3: Process the result
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil,err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil,err
		}
	}

	// Step 4: Print the user details
	fmt.Println("User fetched successfully:", user)

	//return user
	return user,err
}


