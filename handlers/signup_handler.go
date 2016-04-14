package handlers

import (
	"encoding/json"
 	"golang.org/x/oauth2"
 	"net/http"
	"github.com/pivotal-golang/lager"
 	"io/ioutil"

)


var fbConfig = &oauth2.Config{
	ClientID:     "754147878053502", // change this to yours
	ClientSecret: "e303c490d2183d7ef8139a08a34e1eff",
	RedirectURL:  "https://delphos-api-storkoxa.c9users.io/signup/facebook", // change this to your webserver adddress
	Scopes:       []string{"email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://www.facebook.com/dialog/oauth",
		TokenURL: "https://graph.facebook.com/oauth/access_token",
 	},
}


 	

type SignupHandler struct {
	logger lager.Logger
}

func NewSignupHandler(logger lager.Logger) *SignupHandler {
	return &SignupHandler{logger: logger}
}

func (h *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.Session("signup-handler")
 	var result string
 	
 	
 
 	
 	code := r.FormValue("code")
	errorReason := r.FormValue("error_reason")
	
	if errorReason != "" {
		//facebook response request fail.
		result = errorReason
	} else if code == "" {
		//Do the first request to facebook
		url := fbConfig.AuthCodeURL("");
		logger.Info("Login into facebook: " + url)
	
		http.Redirect(w, r, url, http.StatusFound)
  	} else {
		//facebook response success.
		token, err := fbConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			result = "invalid token" 	
		} else {
			//get Data
			
			response, err := http.Get("https://graph.facebook.com/me?fields=name,email&access_token=" + token.AccessToken)
	    	if err != nil {
	    		result = "fail to get the data" 	
	    	 	
	    	} else {
		    	str, err := ioutil.ReadAll(response.Body)
		    	if err != nil {
		    		    		result = "fail to read the body" 	
		    	} else {
	   				result = string(str)
		    	}
	    	}
	  	}
  	}


	infoResponse := map[string]interface{}{
		"facebook": result,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(infoResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("failed-marshaling-signup", err)
		return
	}
}
