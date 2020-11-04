package repositories

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"database/sql"
)

type IFriendRepo interface {
	CreateFriend(*model.FriendsRepoInput) error
	IsBlockedEachOther(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
}

type FriendRepo struct {
	Db *sql.DB
}

func (_self FriendRepo) CreateFriend(friendsRepoInput *model.FriendsRepoInput) error {
	query := `insert into friends(firstid, secondid) values ($1, $2)`
	_, err := _self.Db.Exec(query, friendsRepoInput.FirstID, friendsRepoInput.SecondID)
	return err
}

func (_self FriendRepo) IsBlockedEachOther(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(select true from blocks WHERE (
    						    	requestorid in ($1, $2) 
								    AND 
    						    	targetid in ($1, $2)
    						      ))`
	var isBlocked bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&isBlocked)
	if err != nil {
		return true, err
	}
	if isBlocked {
		return true, nil
	}
	return false, nil
}

func (_self FriendRepo) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(
    						select true 
    						from friends 
    						where (
    						    	firstid in ($1, $2) 
								    AND 
    						    	secondid in ($1, $2)
    						      )
    						)`
	var existed bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&existed)
	if err != nil {
		return true, err
	}
	if existed {
		return true, nil
	}
	return false, nil
}
