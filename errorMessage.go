package qqConnect
var ErrorMessage map[string]map[string]interface{}

func init()  {
	ErrorMessage = map[string]map[string]interface{}{
		"32101"		:	{
			"code"	:	"32101",
			"Msg"	:	"",
		},
		"32102"		:	{
			"code"	:	"32102",
			"Msg"	:	"The state does not match. You may be a victim of CSRF.",
		},
		"32103"		:	{
			"code"	:	"32103",
			"Msg"	:	"",
		},
	}
}
