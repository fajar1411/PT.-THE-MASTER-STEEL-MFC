package repouser

import (
	"errors"
	"master/domain/contract/repocontract"
	"master/domain/model"
	"master/domain/query"
	"master/domain/request"

	"gorm.io/gorm"
)

type RepoUser struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) repocontract.RepoUser {
	return &RepoUser{
		db: db,
	}
}

func (ru *RepoUser) RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error) {
	datareqtomodeluser := query.RequuserToModel(newRequest)

	_, errexist := ru.EmaiuserExist(datareqtomodeluser.Email)

	if errexist == nil {
		return request.RequestUser{}, errors.New("Email Sudah Terdaftar")
	}
	tx := ru.db.Create(&datareqtomodeluser)

	if tx.Error != nil {
		return request.RequestUser{}, tx.Error
	}

	datamodeltoreq := query.ModelusertoReq(&datareqtomodeluser)

	return datamodeltoreq, nil
}

func (ru *RepoUser) AllUser() (data []request.RequestUser, err error) {
	var activ []model.User
	tx := ru.db.Raw("Select users.id, users.password, users.email,users.nama from users").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModeluserToReq(activ)
	return dtmdlttoreq, nil
}

func (ru *RepoUser) EmaiuserExist(email string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.password, users.email, users.nama from users WHERE users.email= ? ", email).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModelusertoReq(&activ)
	return activcore, nil
}

// // Player implements repocontract.RepoUser.
func (rp *RepoUser) Profile(id int) (data request.RequestUser, err error) {
	profile := model.User{}

	tx := rp.db.Raw("Select users.id, users.password, users.email, users.nama from users WHERE users.id= ?", id).Find(&profile)

	if tx.Error != nil {
		return request.RequestUser{}, tx.Error
	}
	list := query.ModelusertoReq(&profile)

	return list, nil
}

// // Filter implements repocontract.RepoUser.
// func (ru *RepoUser) Filter(nama string, norek string) (data []request.ReqProfile, err error) {

// 	fmt.Print("ini nama", nama)
// 	profile := []model.Profile{}

// 	if nama != "" && norek == "" {
// 		// Filter by nama only
// 		tx1 := ru.db.Raw("SELECT users.id, users.email, users.nama, wallets.nama_dompet, wallets.account_wallet, wallets.saldo, banks.no_rekening FROM users LEFT JOIN wallets ON wallets.id_player = users.id LEFT JOIN banks ON banks.id_player = users.id WHERE users.nama LIKE ?", "%"+nama+"%").Find(&profile)

// 		if tx1.Error != nil {
// 			return []request.ReqProfile{}, tx1.Error
// 		}

// 		list := query.FiltemodeluserToReq(profile)
// 		return list, nil
// 	}

// 	if norek != "" && nama == "" {
// 		// Filter by norek only
// 		tx2 := ru.db.Raw("SELECT users.id, users.email, users.nama, wallets.nama_dompet, wallets.account_wallet, wallets.saldo, banks.no_rekening FROM users LEFT JOIN wallets ON wallets.id_player = users.id LEFT JOIN banks ON banks.id_player = users.id WHERE banks.no_rekening =?", norek).Find(&profile)

// 		if tx2.Error != nil {
// 			return []request.ReqProfile{}, tx2.Error
// 		}

// 		list := query.FiltemodeluserToReq(profile)
// 		return list, nil
// 	}

// 	if nama != "" && norek != "" {
// 		// Filter by both nama and norek
// 		tx3 := ru.db.Raw("SELECT users.id, users.email, users.nama, wallets.nama_dompet, wallets.account_wallet, wallets.saldo, banks.no_rekening FROM users LEFT JOIN wallets ON wallets.id_player = users.id LEFT JOIN banks ON banks.id_player = users.id WHERE users.nama LIKE ? AND banks.no_rekening = ?", "%"+nama+"%", norek).Find(&profile)

// 		if tx3.Error != nil {
// 			return []request.ReqProfile{}, tx3.Error
// 		}

// 		list := query.FiltemodeluserToReq(profile)
// 		fmt.Print("ini repo", list)
// 		return list, nil
// 	}
// 	return []request.ReqProfile{}, nil
// }
