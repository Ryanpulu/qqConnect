package qqConnect

import "qqConnect/providers"

type QC struct {
	AppID			string
	AppKey			string
	Scope			string
	Callback		string
	ErrorReport		bool
}

func (self *QC)Login(){
	providers.Test()
}
