package weixin

import (
	"context"
	"fmt"
	"github.com/axengine/utils"
	"github.com/axengine/utils/random"
	"github.com/labstack/echo/v4"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"testing"
	"time"
)

var (
	mchId    = "16--------"
	apiV3Key = "hI----------"
	appId    = "wx-----------------"
	serialNo = "3E---------------"
)

func TestPrePay(t *testing.T) {
	client, err := PayClient(context.Background(), "apiclient_key.pem", mchId, serialNo, apiV3Key)
	if err != nil {
		t.Fatal(err)
	}
	svc := native.NativeApiService{Client: client}
	orderId := random.RandAlphaAndDigits(32, true)
	fmt.Println(orderId)
	resp, _, err := svc.Prepay(context.Background(), native.PrepayRequest{
		Appid:       core.String(appId),
		Mchid:       core.String(mchId),
		Description: core.String("日会员订阅"),
		OutTradeNo:  core.String(orderId),
		TimeExpire:  core.Time(time.Now().Add(time.Minute * 5)),
		NotifyUrl:   core.String("http://111.111.111.111:8080/order/callback"),
		Amount: &native.Amount{
			Total: core.Int64(int64(1)),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrettyToStdout(resp)
	fmt.Println(resp.String())
}

func TestQueryPay(t *testing.T) {
	client, err := PayClient(context.Background(), "apiclient_key.pem", mchId, serialNo, apiV3Key)
	if err != nil {
		t.Fatal(err)
	}
	svc := native.NativeApiService{Client: client}
	resp, _, err := svc.QueryOrderByOutTradeNo(context.Background(), native.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String("9s4wsgllnx3chzzc7c6yzefxtlyfjyvu"),
		Mchid:      core.String(mchId),
	})
	if err != nil {
		t.Fatal(err)
	}
	print(utils.JsonPretty(resp))
}

// WeixinPayNotify 以echo为例， http 句柄
func WeixinPayNotify(c echo.Context) error {
	ctx := c.Request().Context()
	handler, err := NotifyHandle(ctx, "apiclient_key.pem", mchId, serialNo, apiV3Key)
	if err != nil {
		return err
	}

	// 验签
	transaction := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(ctx, c.Request(), transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return err
	}

	// 处理通知内容
	if *transaction.TradeState == "SUCCESS" {
		// do something
	}

	return nil
}
