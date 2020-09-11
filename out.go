package ogjson


type Out struct {
	Fail bool
	Response Response
}

func OutFail(data interface{}, code ResponseFailCode) Out {
	return Out{
		Fail:true,
		Response: Fail(data, code),
	}
}
func OutFailMsg (msg string) Out {
	return Out{
		Fail:true,
		Response: FailMsg(msg),
	}
}

func OutFailCode( code ResponseFailCode) Out {
	return Out{
		Fail:true,
		Response: FailCode(code),
	}
}

func OutAuth (code ResponseAuthCode) Out {
	return Out{
		Fail:true,
		Response: Auth(code),
	}
}