package config

type Payment struct {
	Title  string
	Wechat *WechatPay
	Alipay *Alipay
}

type WechatPay struct {
	Appid       string
	MchId       string
	MchSerialNo string
	ApiKey      string
	ApiV3Key    string
	// CertFileContent
	PrivateKeyContent string
	// Pkcs12FileContent string
	Domain string
}

type Alipay struct {
	Mode                    string
	Appid                   string
	Openid                  string
	PrivateKey              string
	AppPublicCertContent    string
	AlipayRootCertContent   string
	AlipayPublicCertContent string
	ContentKey              string
}
