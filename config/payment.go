package config

type PayPlatform struct {
	Title  string     `mapstructure:"title" json:"title" yaml:"title"`
	Wechat *WechatPay `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
	Alipay *Alipay    `mapstructure:"alipay" json:"alipay" yaml:"alipay"`
}

type WechatPay struct {
	Appid             string `mapstructure:"appid" json:"appid" yaml:"appid"`
	MchId             string `mapstructure:"mchId" json:"mchId" yaml:"mchId"`
	MchSerialNo       string `mapstructure:"mchSerialNo" json:"mchSerialNo" yaml:"mchSerialNo"`
	ApiKey            string `mapstructure:"apiKey" json:"apiKey" yaml:"apiKey"`
	ApiV3Key          string `mapstructure:"apiV3Key" json:"apiV3Key" yaml:"apiV3Key"`
	PrivateKeyContent string `mapstructure:"privateKeyContent" json:"privateKeyContent" yaml:"privateKeyContent"`
}

type Alipay struct {
	Mode                    string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Appid                   string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Openid                  string `mapstructure:"openid" json:"openid" yaml:"openid"`
	PrivateKey              string `mapstructure:"privateKey" json:"privateKey" yaml:"privateKey"`
	AppPublicCertContent    string `mapstructure:"appPublicCertContent" json:"appPublicCertContent" yaml:"appPublicCertContent"`
	AlipayRootCertContent   string `mapstructure:"alipayRootCertContent" json:"alipayRootCertContent" yaml:"alipayRootCertContent"`
	AlipayPublicCertContent string `mapstructure:"alipayPublicCertContent" json:"alipayPublicCertContent" yaml:"alipayPublicCertContent"`
	ContentKey              string `mapstructure:"contentKey" json:"contentKey" yaml:"contentKey"`
}
