package weixin

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func PayClient(ctx context.Context, appClientKeyFile, mchId, certificateSerialNo, mchAPIv3Key string) (*core.Client, error) {
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(appClientKeyFile)
	if err != nil {
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchId,
			certificateSerialNo,
			mchPrivateKey,
			mchAPIv3Key),
	}
	return core.NewClient(ctx, opts...)
}

func NotifyHandle(ctx context.Context, appClientKeyFile, mchId, certificateSerialNo, mchAPIv3Key string) (*notify.Handler, error) {
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(appClientKeyFile)
	if err != nil {
		return nil, err
	}
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	if err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey,
		certificateSerialNo, mchId, mchAPIv3Key); err != nil {
		return nil, err
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchId)

	// 3. 使用证书访问器初始化 `notify.Handler`
	handler, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	return handler, err
}
