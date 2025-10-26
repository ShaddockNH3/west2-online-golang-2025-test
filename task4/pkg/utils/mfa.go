package utils

import (
	"encoding/base64"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func GenerateMFAData(username string) (string, string, error) {
	// 直接用 totp.Generate 生成密钥和 URL，它内部会使用安全的随机数
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      constants.AppName,
		AccountName: username,
	})
	if err != nil {
		return "", "", err
	}

	// 把 key.URL() 画成二维码
	pngBytes, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		return "", "", err
	}

	// 转换成 base64 返回给前端
	qrCodeBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(pngBytes)

	// 返回秘密密钥和二维码数据
	return key.Secret(), qrCodeBase64, nil
}
