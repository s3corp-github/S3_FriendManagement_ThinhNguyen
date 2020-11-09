package repositories

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"database/sql"
	"fmt"
	"strings"
)

type IUserRepo interface {
	CreateUser(*model.UserRepoInput) error
	IsExistedUser(string) (bool, error)
	GetUserIDByEmail(string) (int, error)
	GetEmailListByIDs(userIDs []int) ([]string, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (_self UserRepo) CreateUser(userRepoInput *model.UserRepoInput) error {
	query := `insert into useremails(email) values ($1)`
	_, err := _self.Db.Exec(query, userRepoInput.Email)
	return err
}

func (_self UserRepo) GetUserIDByEmail(email string) (int, error) {
	query := `select id from useremails where email=$1`
	var userID int
	err := _self.Db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return userID, nil
}

func (_self UserRepo) IsExistedUser(email string) (bool, error) {
	query := `select exists (select true from useremails where email=$1)`
	var existed bool
	err := _self.Db.QueryRow(query, email).Scan(&existed)
	if err != nil {
		return false, err
	}
	if existed {
		return true, nil
	}
	return false, nil
}

func (_self UserRepo) GetEmailListByIDs(userIDs []int) ([]string, error) {
	if len(userIDs) == 0 {
		return []string{}, nil
	}

	IDList := make([]string, len(userIDs))
	for i, id := range userIDs {
		IDList[i] = fmt.Sprintf("%v", id)
	}
	query := fmt.Sprintf(`select email from useremails where id in (%v)`, strings.Join(IDList, ","))
	rows, err := _self.Db.Query(query)
	if err != nil {
		return nil, err
	}

	emailList := make([]string, 0)
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emailList = append(emailList, email)
	}
	return emailList, nil
}
