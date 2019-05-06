package app

import (
	"github.com/m-lukas/github-analyser/controller"

	"bitbucket.org/timbertom/backend/httputil"
	"bitbucket.org/timbertom/backend/translate"
)

type scoreResponse struct {
	ActivityScore   float64
	PopularityScore float64
}

func doGetScore(userName string) (*scoreResponse, *httputil.ErrorResponse) {

	if userName == "" {
		return nil, httputil.FromTranslationKey(400, translate.MissingParameter)
	}

	user, err := controller.GetUser(userName)
	if err != nil {
		return nil, httputil.FromTranslationKey(500, translate.ServerError)
	}

	resp := &scoreResponse{
		ActivityScore:   user.ActivityScore,
		PopularityScore: user.PopularityScore,
	}

	return resp, nil

}
