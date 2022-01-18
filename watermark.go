package codecprofile

import (
	"errors"
	"fmt"
)

type Watermark struct {
	// ID uid_name
	ID string `json:"_id" json:"id"`

	// Uid 用户uid
	Uid int `bson:"uid" json:"uid"`

	// Name 水印名称
	Name string `bson:"name" json:"name"`

	// Desc 水印描述：对水印模版对文字描述
	Desc string `bson:"desc,omitempty" json:"desc,omitempty"`

	// Profile 水印配置
	Profile WatermarkProfile `bson:"profile" json:"profile"`

	// CustomProfile 自定义转码配置
	CustomProfile string `bson:"customProfile,omitempty" json:"customProfile,omitempty"`
}

type WatermarkProfile struct {
	// Image 图片：水印图片信息, 不为空
	Image string `bson:"image" json:"image"`

	// ImageHigh 图片原始高度：水印图片原始高度，数字，单位px 大于0
	ImageHigh float64 `bson:"imageHigh" json:"imageHigh"`

	// ImageWidth 图片原始宽度：水印图片原始宽度，数字，单位px 大于0
	ImageWidth float64 `bson:"imageWidth" json:"imageWidth"`

	// X 横坐标：水印图片左边距视频左边占视频总宽度的比例，取值范围[0,1]
	X float64 `bson:"x" json:"x"`

	// Y 纵坐标：水印图片上边距视频上边占视频总高度的比例，取值范围[0,1]
	Y float64 `bson:"y" json:"y"`

	// WidthProportion 宽度比例：水印图片宽度占视频宽度的比例，取值范围[0,1]
	WidthProportion float64 `bson:"widthProportion" json:"widthProportion"`
}

func (wm *Watermark) IsValid() error {
	if len(wm.Name) == 0 {
		return errors.New("watermark name is invalid")
	}
	if wm.Uid <= 0 {
		return fmt.Errorf("watermark uid is invalid, uid: %v", wm.Uid)
	}
	if len(wm.Profile.Image) == 0 {
		return fmt.Errorf("watermark image is invalid, image: %v", wm.Profile.Image)
	}
	if wm.Profile.ImageHigh <= 0 {
		return fmt.Errorf("watermark ImageHigh is invalid, ImageHigh: %v", wm.Profile.ImageHigh)
	}
	if wm.Profile.ImageWidth <= 0 {
		return fmt.Errorf("watermark ImageWidth is invalid, ImageWidth: %v", wm.Profile.ImageWidth)
	}
	if wm.Profile.X < 0 || wm.Profile.X > 1 {
		return fmt.Errorf("watermark x is invalid, x: %v", wm.Profile.X)
	}
	if wm.Profile.Y < 0 || wm.Profile.Y > 1 {
		return fmt.Errorf("watermark y is invalid, y: %v", wm.Profile.Y)
	}
	if wm.Profile.WidthProportion < 0 || wm.Profile.WidthProportion > 1 {
		return fmt.Errorf("watermark WidthProportion is invalid, WidthProportion: %v", wm.Profile.WidthProportion)
	}
	return nil
}
