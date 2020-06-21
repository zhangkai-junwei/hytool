package qrocde

import (
	"github.com/makiuchi-d/gozxing"
	gq "github.com/makiuchi-d/gozxing/qrcode"
	"github.com/outakujo/utils"
	"github.com/skip2/go-qrcode"
	"image"
	"os"
)

type MQr struct {
}

func (*MQr) Gen(data, file string) {
	code, e := qrcode.New(data, qrcode.Medium)
	utils.PanicError(e)
	code.DisableBorder = true
	e = code.WriteFile(256, file)
	utils.PanicError(e)
}

func (*MQr) Parse(file string) (string, error) {
	fi, _ := os.Open(file)
	img, _, _ := image.Decode(fi)
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := gq.NewQRCodeReader()
	result, e := qrReader.Decode(bmp, nil)
	if e != nil {
		return "", e
	}
	return result.String(), e

}
