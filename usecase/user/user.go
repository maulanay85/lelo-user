package user

import (
	"context"
	"fmt"

	entity "lelo-user/entity"
	util "lelo-user/util"
	utilConstants "lelo-user/util/constants"

	log "github.com/sirupsen/logrus"
)

func (u *UserUsecaseModule) RegisterUser(ctx context.Context, fullname string, email string, pass string) (int64, error) {
	log.Infof("RegisterUser for email: %s", email)
	id, err := u.Repo.CheckEmailExist(ctx, email)
	if err != nil {
		log.Errorf("[usecase] RegisterUser: %v", err)
		errorWrap := fmt.Errorf("error CheckEmailExist: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	if id != 0 {
		log.Error("[usecase] RegisterUser: email is exist")
		errorWrap := fmt.Errorf("email is exist: %w", util.ErrorPreCondition)
		return 0, errorWrap
	}
	// hash password
	hash, err := u.UtilAuth.HashPassword(pass)
	if err != nil {
		log.Errorf("[usecase] RegisterUser.HashPassword: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}

	tx, err := u.UtilDbModule.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	if err != nil {
		log.Errorf("[usecase] RegisterUser.BeginTx: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}

	// get role
	role, err := u.RoleRepository.GetRoleByCode(ctx, tx, utilConstants.USER)
	if err != nil {
		log.Errorf("[usecase] RegisterUser.GetRoleByCode: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}

	// insert user
	user := entity.UserEntity{
		Email:    email,
		Fullname: fullname,
		Pass:     hash,
	}
	iduser, err := u.Repo.InsertUserTx(ctx, tx, &user)
	if err != nil {
		log.Errorf("[usecase] RegisterUser.InsertUserTx: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}

	// inser map user role
	userRole := entity.UserRoleEntity{
		UserId: iduser,
		RoleId: role.Id,
	}
	_, err = u.UserRuleRepo.InsertUserRoleTx(ctx, tx, &userRole)

	if err != nil {
		log.Errorf("[usecase] RegisterUser.InsertUserRoleTx: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}

	log.Infof("Finish RegisterUser for email: %s", email)
	return iduser, nil
}

func (u *UserUsecaseModule) Login(ctx context.Context, email string, pass string) (*entity.TokenEntity, error) {
	user, err := u.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Errorf("[usecase] Login: %v", err)
		errorWrap := fmt.Errorf("wrong email or pass : %w", util.ErrorUnauthenticated)
		return nil, errorWrap
	}
	if isValid := u.UtilAuth.CheckHashPassword(pass, user.Pass); !isValid {
		errorWrap := fmt.Errorf("wrong email or pass :%w", util.ErrorUnauthenticated)
		return nil, errorWrap
	}
	userRole, err := u.UserRuleRepo.GetUserRoleByUserId(ctx, user.Id)
	if err != nil {
		log.Errorf("[usecase] Login: %v", err)
		errorWrap := fmt.Errorf("error :%w", util.ErrorInternalServer)
		return nil, errorWrap
	}

	token, err := u.UtilAuth.GenerateToken(userRole)
	if err != nil {
		log.Errorf("[usecase] GenerateToken: %v", err)
		errorWrap := fmt.Errorf("error :%w", util.ErrorInternalServer)
		return nil, errorWrap
	}
	return token, nil
}
