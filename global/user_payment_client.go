package global

import (
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat/v3"
)

type UserPaymentClient struct {
	// rwMu   sync.RWMutex
	Alipay *alipay.Client
	Wxpay  *wechat.ClientV3
}

func NewUserPaymentService() {
	// 对加密后的配置内容进行解密
	appid, err := DecryptByAes(GVA_CONFIG.PayPlatform.Alipay.Appid)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Alipay.Appid = string(appid)

	openid, err := DecryptByAes(GVA_CONFIG.PayPlatform.Alipay.Openid)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Alipay.Openid = string(openid)

	privateKey, err := DecryptByAes(GVA_CONFIG.PayPlatform.Alipay.PrivateKey)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Alipay.PrivateKey = string(privateKey)

	contentKey, err := DecryptByAes(GVA_CONFIG.PayPlatform.Alipay.ContentKey)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Alipay.ContentKey = string(contentKey)

	// alipay client
	isPord := true
	if GVA_CONFIG.PayPlatform.Alipay.Mode != "pord" {
		isPord = false
	}
	alipayCli, err := alipay.NewClient(GVA_CONFIG.PayPlatform.Alipay.Appid, GVA_CONFIG.PayPlatform.Alipay.PrivateKey, isPord)
	if err != nil {
		GVA_LOG.Error(err.Error())
		// panic(err)
	}
	// Debug开关，输出/关闭日志
	alipayCli.DebugSwitch = gopay.DebugOff
	// 配置公共参数
	alipayCli.SetCharset(alipay.UTF8).
		SetSignType(alipay.RSA2).
		// SetAppAuthToken("").
		// SetReturnUrl(GVA_CONFIG.PayPlatform.Alipay.ReturnUrl).
		SetNotifyUrl(fmt.Sprintf("%s%s", GVA_CONFIG.System.Domain, "/api/v1/mall/asset/deposit/alipay/notify"))
		// SetAESKey(GVA_CONFIG.PayPlatform.Alipay.ContentKey)

	// 自动同步验签（只支持证书模式）
	// 传入 支付宝公钥证书 alipayPublicCert.crt 内容
	alipayCli.AutoVerifySign([]byte(GVA_CONFIG.PayPlatform.Alipay.AlipayPublicCertContent))
	// 传入证书内容
	if err = alipayCli.SetCertSnByContent([]byte(GVA_CONFIG.PayPlatform.Alipay.AppPublicCertContent),
		[]byte(GVA_CONFIG.PayPlatform.Alipay.AlipayRootCertContent),
		[]byte(GVA_CONFIG.PayPlatform.Alipay.AlipayPublicCertContent)); err != nil {
		GVA_LOG.Error(err.Error())
		// panic(err)
	}

	// 对加密后的配置内容进行解密
	appid1, err := DecryptByAes(GVA_CONFIG.PayPlatform.Wechat.Appid)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Wechat.Appid = string(appid1)

	mchid, err := DecryptByAes(GVA_CONFIG.PayPlatform.Wechat.MchId)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Wechat.MchId = string(mchid)

	mchSerialNo, err := DecryptByAes(GVA_CONFIG.PayPlatform.Wechat.MchSerialNo)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Wechat.MchSerialNo = string(mchSerialNo)

	apiV3Key, err := DecryptByAes(GVA_CONFIG.PayPlatform.Wechat.ApiV3Key)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Wechat.ApiV3Key = string(apiV3Key)

	privateKeyContent, err := DecryptByAes(GVA_CONFIG.PayPlatform.Wechat.PrivateKeyContent)
	if err != nil {
		GVA_LOG.Error(err.Error())
	}
	GVA_CONFIG.PayPlatform.Wechat.PrivateKeyContent = string(privateKeyContent)

	// NewClientV3 初始化微信客户端 V3
	//
	//	mchid：商户ID
	//	serialNo：商户证书的证书序列号
	//	apiV3Key：APIv3Key，商户平台获取
	//	privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	wechatCli, err := wechat.NewClientV3(GVA_CONFIG.PayPlatform.Wechat.MchId, GVA_CONFIG.PayPlatform.Wechat.MchSerialNo,
		GVA_CONFIG.PayPlatform.Wechat.ApiV3Key, GVA_CONFIG.PayPlatform.Wechat.PrivateKeyContent)
	if err != nil {
		GVA_LOG.Error(err.Error())
		// panic(err)
	}

	// 设置微信平台证书和序列号，如开启自动验签，请忽略此步骤
	// wechatCli.SetPlatformCert([]byte(GVA_CONFIG.PayPlatform.Wechat.),)

	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = wechatCli.AutoVerifySign(true)
	if err != nil {
		GVA_LOG.Error(err.Error())
		// panic(err)
	}

	// 打开Debug开关，输出日志
	wechatCli.DebugSwitch = gopay.DebugOff

	PaymentClient = &UserPaymentClient{
		Alipay: alipayCli,
		Wxpay:  wechatCli,
	}
}
