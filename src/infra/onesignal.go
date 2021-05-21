package infra

type OneSignal struct {
	AppId  string
	AppKey string
}

func LoadOnesignalConfig() OneSignal {
	return OneSignal{
		AppId:  RequireEnvString("ONESIGNAL_APPID"),
		AppKey: RequireEnvString("ONESIGNAL_APPKEY"),
	}
}
