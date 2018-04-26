package qqConnect

import (
	"io/ioutil"
	"encoding/json"
)
var OauthApiMothod map[string]interface{}
type QC struct {
	AppID			string
	AppKey			string
	OpenID			string
	AccessToken		string
	Scope			string
	CallbackUri		string
	ErrorReport		bool
}

func (self *QC)Login(){

}

func (self *QC)Callback(){

}

func init(){
	readOauthApiMethod()
}

func (self *QC)GetDataByOauth(method string){

}

func readOauthApiMethod(){
	fileByte,err := ioutil.ReadFile("config/oauthApiMethod.json");
	if nil != err {
		panic(err)
	}
	json.Unmarshal(fileByte,&OauthApiMothod)
}
