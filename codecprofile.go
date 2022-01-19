package codecprofile

import (
	"fmt"
	"strconv"
)

const (
	CodecTypeNormal = ""
	CodecTypeSmart  = "smart"
)

const (
	CodecStatusOpen  = "open"
	CodecStatusClose = "close"
)

type CodecInfo struct {
	// ID 唯一ID uid_name
	ID string `bson:"_id" json:"id"`

	// Name 转码配置名, 大小写英文、数字、中划线和下划线, 长度[4,6stringd下唯一
	Name string `bson:"name" json:"name"`

	// Uid 用户UID，预置模板不作设置
	Uid int64 `bson:"uid" json:"uid"`

	// Status 发布状态，仅对预置模板有效，open,close
	Status string `bson:"status,omitempty" json:"status,omitempty"`

	// CodecType 转码类型 空:标准类型，smart:锐智转码
	CodecType string `bson:"codecType,omitempty" json:"codecType,omitempty"`

	// WatermarkID 水印模板ID
	WatermarkID string `bson:"watermarkID,omitempty" json:"watermarkID,omitempty"`

	// Profile 转码配置
	Profile CodecProfile `bson:"profile,omitempty" json:"profile,omitempty"`

	// CustomProfile 自定义转码配置
	CustomProfile string `bson:"customProfile,omitempty" json:"customProfile,omitempty"`
}

type CodecProfile struct {
	// Vcodec 使用的视频编码器，如libx264、libx265，在不改变编码时为copy；默认值为copy（原视频编码)
	Vcodec string `bson:"vcodec,omitempty" json:"vcodec,omitempty"`

	// VideoLongSide 视频长边 单位px
	VideoLongSide int `bson:"videoLongSide,omitempty" json:"videoLongSide,omitempty"`

	// VideoShortSide 视频短边 单位px
	VideoShortSide int `bson:"videoShortSide,omitempty" json:"videoShortSide,omitempty"`

	// VideoWidth 视频宽度 单位px
	VideoWidth int `bson:"videoWidth,omitempty" json:"videoWidth,omitempty"`

	// VideoHigh 视频高度 单位px
	VideoHigh int `bson:"videoHigh,omitempty" json:"videoHigh,omitempty"`

	// VideoBitRate 视频码率 单位：千比特每秒（kbit/s）。默认为0，不强制设置码率（自动）。
	// 在不改变视频编码格式时 若指定码率大于原视频码率，则使用原视频码率进行转码。码率控制方式默认为vbr
	VideoBitRate int `bson:"videoBitRate,omitempty" json:"videoBitRate,omitempty"`

	// VideoFrameRate 视频帧率 r，数字。常用帧率24、25、30、60。默认为0，不改变原视频帧率
	VideoFrameRate int `bson:"videoFrameRate,omitempty" json:"videoFrameRate,omitempty"`

	// KeyframeInterval 关键帧间隔 force_key_frames，数字，单位秒。默认为0，不改变原视频关键帧间隔
	KeyframeInterval int `bson:"keyframeInterval,omitempty" json:"keyframeInterval,omitempty"`

	// RemoveBFrame 去B帧 默认关闭
	RemoveBFrame bool `bson:"removeBFrame,omitempty" json:"removeBFrame,omitempty"`

	// KeepSei 保留SEI 默认关闭
	KeepSei bool `bson:"keepSei,omitempty" json:"keepSei,omitempty"`

	// AudioVideoType 音视频输出 0.同时输出音视频，1.只输出视频，2.只输出音频，默认为0。
	// 当设置为只输出视频时，音频相关参数失效；当只输出音频时，视频相关设置失效
	AudioVideoType int `bson:"audioVideoType,omitempty" json:"audioVideoType,omitempty"`

	// Acodec 音频编码 acodec，使用的音频编码器，如aac、opus，在不改变编码时为copy。默认值为copy（原音频编码）
	Acodec string `bson:"acodec,omitempty" json:"acodec,omitempty"`

	// Ac 声道数 ac，数字。默认为空，不改变原视频声道数
	Ac int `bson:"ac,omitempty" json:"ac,omitempty"`

	// Ab 音频码率 ab，数字，单位：千比特每秒（kbit/s），常用码率：64k，128k，192k，256k，320k等。默认为0，不强制设置码率（自动）。
	// 在不改变音频编码格式时，若指定码率大于原音频码率，则使用原音频码率进行转码
	Ab int `bson:"ab,omitempty" json:"ab,omitempty"`

	// Ar 音频采样率 ar，数字，单位：赫兹（Hz）。常用音频采样频率有 8000, 11025、12000、16000、22050、24000、32000、44056、44100、47250、48000、50000、64000、88200、96000 等。
	// 默认为0，不改变原视频的音频采样率
	Ar int `bson:"ar,omitempty" json:"ar,omitempty"`

	// BufSize 视频码率控制缓冲区大小 bufsize，数字，单位：千比特每秒（kbit/s）。默认为空
	BufSize int `bson:"bufSize,omitempty" json:"bufSize,omitempty"`

	// MaxRate 最大视频码率容忍度 maxrate，数字，单位：千比特每秒（kbit/s）。默认为空
	MaxRate int `bson:"maxRate,omitempty" json:"maxRate,omitempty"`

	// MinRate 最小视频码率容忍度 minrate，数字，单位：千比特每秒（kbit/s）。默认为空
	MinRate int `bson:"minRate,omitempty" json:"minRate,omitempty"`

	// PixFmt 像素格式 pix_fmt，如yuv420p。默认为空
	PixFmt string `bson:"pixFmt,omitempty" json:"pixFmt,omitempty"`

	// Crf，数字，取值范围为 0-51，值越小品质越高，0 为无损。默认为空
	Crf int `bson:"crf,omitempty" json:"crf,omitempty"`

	// Preset 影响文件的生成速度，速度越快，压缩比率越低，有以下几个可用的值 ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow。默认为空
	Preset string `bson:"preset,omitempty" json:"preset,omitempty"`
}

func (ci *CodecInfo) IsValid() error {
	if len(ci.Name) == 0 {
		return fmt.Errorf("CodecInfo name is invalid, name: %v", ci.Name)
	}
	if ci.Uid < 0 {
		return fmt.Errorf("CodecInfo uid is invalid")
	}
	if ci.Status != CodecStatusOpen && ci.Status != CodecStatusClose {
		return fmt.Errorf("CodecInfo Status is invalid, Status: %v", ci.Status)
	}
	if ci.CodecType != CodecTypeNormal && ci.Status != CodecTypeSmart {
		return fmt.Errorf("CodecInfo CodecType is invalid, CodecType: %v", ci.CodecType)
	}
	return nil
}

func TransformVideo(codecProfile CodecProfile) (args map[string]string, err error) {
	if codecProfile.VideoLongSide < 0 || codecProfile.VideoShortSide < 0 || codecProfile.VideoWidth < 0 || codecProfile.VideoHigh < 0 ||
		codecProfile.VideoBitRate < 0 || codecProfile.VideoFrameRate < 0 || codecProfile.KeyframeInterval < 0 ||
		codecProfile.BufSize < 0 || codecProfile.MaxRate < 0 || codecProfile.MinRate < 0 || codecProfile.Crf < 0 {
		err = fmt.Errorf("invalid video codecProfile")
		return nil, err
	}

	videocodecargs := make(map[string]string)

	// 设置视频编码器
	vcodec := "copy"
	if codecProfile.Vcodec != "" {
		vcodec = codecProfile.Vcodec
	}
	videocodecargs["-vcodec"] = vcodec

	useside := true
	if codecProfile.VideoLongSide == 0 && codecProfile.VideoShortSide == 0 {
		useside = false
	}

	scale := ""
	if useside { // 使用长短边
		longside := ""
		shortside := ""
		uselongside := false
		useshortside := false
		if codecProfile.VideoLongSide > 0 {
			if codecProfile.VideoShortSide > 0 { //当长短边都指定时，需要按原视频长短边，强制缩放至目标分辨率
				longside = strconv.Itoa(codecProfile.VideoLongSide)
				shortside = strconv.Itoa(codecProfile.VideoShortSide)
			} else { //当只有长边指定时，以这条边为基准，按原视频比例自适应
				uselongside = true
				longside = strconv.Itoa(codecProfile.VideoLongSide)
			}
		} else {
			if codecProfile.VideoShortSide > 0 { //当只有短边指定时，以这条边为基准，按原视频比例自适应
				useshortside = true
				shortside = strconv.Itoa(codecProfile.VideoShortSide)
			}
		}

		if uselongside { //若指定视频长边大于原视频长边，则使用原视频长边进行转码
			scale = fmt.Sprintf("scale='if(gt(%s, max(iw\\, ih)), max(iw\\, ih), 2*floor(%s/max(iw\\, ih)*iw/2))':'if(gt(%s, max(iw\\, ih)), min(iw\\, ih), 2*floor(%s/max(iw\\, ih)*ih/2))'", longside, longside, longside, longside)
		} else if useshortside { //若指定视频短边大于原视频短边，则使用原视频短边进行转码
			scale = fmt.Sprintf("scale='if(gt(%s, min(iw\\, ih)), max(iw\\, ih), 2*floor(%s/min(iw\\, ih)*iw/2))':'if(gt(%s, min(iw\\, ih)), max(iw\\, ih), 2*floor(%s/min(iw\\, ih)*ih/2))'", shortside, shortside, shortside, shortside)
		} else {
			scale = fmt.Sprintf("scale='if(gt(2*floor(%s*max(iw\\,ih)/(min(iw\\,ih)*2))\\,%s)\\,2*floor(%s*(iw/2)/max(iw\\,ih))\\,2*floor(%s*(iw/2)/min(iw\\,ih)))':'if(gt(2*floor(%s*max(iw\\,ih)/(min(iw\\,ih)*2))\\,%s)\\,2*floor(%s*(ih/2)/max(iw\\,ih))\\,2*floor(%s*(ih/2)/min(iw\\,ih)))'", shortside, longside, longside, shortside, shortside, longside, longside, shortside)
		}
	} else { // 使用分辨率
		width := ""
		high := ""

		if codecProfile.VideoWidth > 0 {
			width = strconv.Itoa(codecProfile.VideoWidth)
			if codecProfile.VideoHigh > 0 {
				high = strconv.Itoa(codecProfile.VideoHigh)
			} else {
				high = fmt.Sprintf("%d*ih/iw", codecProfile.VideoWidth)
			}
		} else {
			if codecProfile.VideoHigh > 0 {
				width = fmt.Sprintf("%d*iw/ih", codecProfile.VideoHigh)
				high = strconv.Itoa(codecProfile.VideoHigh)
			}
		}

		if width != "" && high != "" {
			scale = fmt.Sprintf("scale='if(gt(%s,iw), iw\\, %s)':'if(gt(%s,ih)\\, ih\\, %s)'", width, width, high, high)
		}
	}

	if scale != "" {
		if videocodecargs["-vcodec"] == "copy" { // 修改分辨率需要编码器重新编码
			err = fmt.Errorf("invalid Vcodec")
			return nil, err
		}
		videocodecargs["-vf"] = scale
	}

	// 设置视频码率
	if codecProfile.VideoBitRate > 0 {
		videocodecargs["-b:v"] = fmt.Sprintf("%dk", codecProfile.VideoBitRate)
	}

	// 设置视频帧率
	if codecProfile.VideoFrameRate > 0 {
		videocodecargs["-r"] = strconv.Itoa(codecProfile.VideoFrameRate)
	}

	// 设置关键帧间隔
	if codecProfile.KeyframeInterval > 0 {
		videocodecargs["-force_key_frames"] = fmt.Sprintf("\"expr:gte(t,n_forced*%d)\"", codecProfile.KeyframeInterval)
	}

	// 设置去除B帧标志
	if codecProfile.RemoveBFrame {
		bframe := "\"bframes=0\""
		if codecProfile.Vcodec == "libx264" {
			videocodecargs["-x264opts"] = bframe
		} else if codecProfile.Vcodec == "libx265" {
			videocodecargs["-x265-params"] = bframe
		}
	}

	// 设置保留SEI
	if codecProfile.KeepSei {
		// FIXME:暂时忽略
	}

	// 设置视频码率控制缓冲区大小
	if codecProfile.BufSize > 0 {
		videocodecargs["-bufsize"] = fmt.Sprintf("%dk", codecProfile.BufSize)
	}

	// 设置最大视频码率容忍度
	if codecProfile.MaxRate > 0 {
		videocodecargs["-maxrate"] = fmt.Sprintf("%dk", codecProfile.MaxRate)
	}

	// 设置最小视频码率容忍度
	if codecProfile.MinRate > 0 {
		videocodecargs["-minrate"] = fmt.Sprintf("%dk", codecProfile.MinRate)
	}

	// 设置像素格式
	if codecProfile.PixFmt != "" {
		videocodecargs["-pix_fmt"] = codecProfile.PixFmt
	}

	// 设置Crf
	if codecProfile.Crf > 0 {
		videocodecargs["-crf"] = strconv.Itoa(codecProfile.Crf)
	}

	// 设置Preset
	if codecProfile.Preset != "" {
		videocodecargs["-preset"] = codecProfile.Preset
	}

	return videocodecargs, nil
}

func TransformAudio(codecProfile CodecProfile) (args map[string]string, err error) {
	if codecProfile.Ac < 0 || codecProfile.Ab < 0 || codecProfile.Ar < 0 {
		err = fmt.Errorf("invalid audio codecProfile")
		return nil, err
	}

	audiocodecargs := make(map[string]string)

	// 设置音频编码器
	acodec := "copy"
	if codecProfile.Acodec != "" {
		acodec = codecProfile.Acodec
	}
	audiocodecargs["-acodec"] = acodec

	// 设置音频声道数
	if codecProfile.Ac > 0 {
		audiocodecargs["-ac"] = strconv.Itoa(codecProfile.Ac)
	}

	// 设置音频码率
	if codecProfile.Ab > 0 {
		audiocodecargs["-b:a"] = fmt.Sprintf("%dk", codecProfile.Ab)
	}

	// 设置音频采样率
	if codecProfile.Ar > 0 {
		audiocodecargs["-ar"] = strconv.Itoa(codecProfile.Ar)
	}

	return audiocodecargs, nil
}

func Transformwatermark(watermarkProfile WatermarkProfile, videoargs map[string]string) (args map[string]string, err error) {
	watercodecargs := make(map[string]string)
	vf := ""
	scale := ""

	// 取出视频的分辨率用于后期组装-vf
	if _, ok := videoargs["-vf"]; ok {
		scale = videoargs["-vf"]
	}

	// 设置水印信息
	if watermarkProfile.Image != "" {
		if videoargs["-vcodec"] == "copy" {
			// 水印功能需要编码器重新编码
			err = fmt.Errorf("invalid Vcodec")
			return nil, err
		}

		if watermarkProfile.ImageHigh <= 0 || watermarkProfile.ImageWidth <= 0 || watermarkProfile.X < 0 || watermarkProfile.X > 1 ||
			watermarkProfile.Y < 0 || watermarkProfile.Y > 1 || watermarkProfile.WidthProportion < 0 || watermarkProfile.WidthProportion > 1 {
			err = fmt.Errorf("invalid watermarkProfile")
			return nil, err
		}

		// 设置水印图片名
		imagename := fmt.Sprintf("movie=%s[wmo];", watermarkProfile.Image)

		// 设置水印宽高信息
		scale2ref := fmt.Sprintf("[wmo][in]scale2ref=(main_w*%f):((%f/%f)*(main_w*%f))[wm][ino];", watermarkProfile.WidthProportion, watermarkProfile.ImageHigh, watermarkProfile.ImageWidth, watermarkProfile.WidthProportion)

		// 设置水印x和y位置
		overlay := fmt.Sprintf("[ino][wm]overlay=%f*main_w:%f*main_h", watermarkProfile.X, watermarkProfile.Y)

		if scale != "" { // 设置水印和视频分辨率
			vf = fmt.Sprintf("\"%s%s%s[wmimage];[wmimage]%s[out]\"", imagename, scale2ref, overlay, scale)
		} else { //设置水印
			vf = fmt.Sprintf("\"%s%s%s[out]\"", imagename, scale2ref, overlay)
		}
	} else {
		if scale != "" {
			vf = fmt.Sprintf("\"%s\"", scale)
		}
	}

	if vf != "" {
		watercodecargs["-vf"] = vf
	}

	return watercodecargs, nil
}

func Transform(codecProfile CodecProfile, watermarkProfile WatermarkProfile) (map[string]string, error) {
	var err error
	codecargs := make(map[string]string)
	videocodecargs := make(map[string]string)
	audiocodecargs := make(map[string]string)
	watercodecargs := make(map[string]string)

	switch codecProfile.AudioVideoType {
	case 0: //输出音视频
		if videocodecargs, err = TransformVideo(codecProfile); err != nil {
			return nil, err
		}

		if audiocodecargs, err = TransformAudio(codecProfile); err != nil {
			return nil, err
		}
	case 1: // 只输出视频
		// 去除音频
		codecargs["-an"] = ""
		if videocodecargs, err = TransformVideo(codecProfile); err != nil {
			return nil, err
		}
	case 2: //只输出音频
		// 去除视频
		codecargs["-vn"] = ""
		if audiocodecargs, err = TransformAudio(codecProfile); err != nil {
			return nil, err
		}
	default:
		err = fmt.Errorf("invalid AudioVideoType")
		return nil, err
	}

	if watercodecargs, err = Transformwatermark(watermarkProfile, videocodecargs); err != nil {
		return nil, err
	}

	for key, value := range videocodecargs {
		codecargs[key] = value
	}

	for key, value := range audiocodecargs {
		codecargs[key] = value
	}

	for key, value := range watercodecargs {
		codecargs[key] = value
	}

	return codecargs, nil
}
