package managers

import (
	"github.com/alganbr/kedai-authsvc/internal/models"
	"github.com/alganbr/kedai-authsvc/internal/repos"
	"github.com/alganbr/kedai-usersvc-client/client"
	userModels "github.com/alganbr/kedai-usersvc-client/models"
	"github.com/alganbr/kedai-utils/datetime"
	"github.com/alganbr/kedai-utils/errors"
	"net/http"
)

type IAuthManager interface {
	Get(string) (*models.AccessToken, *errors.Error)
	Authenticate(*models.AccessTokenRq) (*models.AccessToken, *errors.Error)
}

type AuthManager struct {
	authRepo      repos.IAuthRepository
	userSvcClient client.IUserSvcClient
}

func NewAuthManager(authRepo repos.IAuthRepository, userSvcClient client.IUserSvcClient) IAuthManager {
	return &AuthManager{
		authRepo:      authRepo,
		userSvcClient: userSvcClient,
	}
}

func (mgr *AuthManager) Get(id string) (*models.AccessToken, *errors.Error) {
	token, err := mgr.authRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (mgr *AuthManager) Authenticate(rq *models.AccessTokenRq) (*models.AccessToken, *errors.Error) {
	if validateErr := rq.Validate(); validateErr != nil {
		return nil, validateErr
	}

	var accessToken *models.AccessToken
	var err *errors.Error

	switch rq.GrantType {
	case "password":
		accessToken, err = mgr.authenticatePassword(rq)
		break
	case "oauth":
		accessToken, err = mgr.authenticateOauth(rq)
		break
	default:
		break
	}

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (mgr *AuthManager) authenticatePassword(rq *models.AccessTokenRq) (*models.AccessToken, *errors.Error) {
	user, err := mgr.userSvcClient.Password().Validate(&userModels.ValidatePasswordRq{
		Email:    rq.Email,
		Password: rq.Password,
	})
	if err != nil {
		return nil, err
	}
	accessToken := &models.AccessToken{
		UserId:  user.Id,
		Expires: datetime.GetUtcNow().AddDate(0, 0, 1).Unix(),
	}
	err = mgr.authRepo.Create(accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (mgr *AuthManager) authenticateOauth(rq *models.AccessTokenRq) (*models.AccessToken, *errors.Error) {
	return nil, &errors.Error{
		Code:    http.StatusNotImplemented,
		Message: "Not implemented",
	}
}
