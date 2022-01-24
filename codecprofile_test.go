package codecprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	at := assert.New(t)
	{
		codec_profile := CodecProfile{}
		watermark := WatermarkProfile{}
		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)
		at.Equal(len(profile), 2)
		at.Equal(profile["-vcodec"], "copy")
		at.Equal(profile["-acodec"], "copy")
	}
	{
		codec_profile := CodecProfile{}
		watermark := WatermarkProfile{
			Image:           "water.png",
			ImageHigh:       320,
			ImageWidth:      240,
			X:               0.5,
			Y:               0.5,
			WidthProportion: 0.1,
		}

		_, err := Transform(codec_profile, watermark)
		at.NotNil(err)
	}
	{
		codec_profile := CodecProfile{
			VideoFrameRate:   -1,
			KeyframeInterval: 10,
			AudioVideoType:   0,
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		_, err := Transform(codec_profile, watermark)
		at.NotNil(err)
	}
	{
		codec_profile := CodecProfile{
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   1,
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		_, err := Transform(codec_profile, watermark)
		at.NotNil(err)
	}
	{
		codec_profile := CodecProfile{
			Vcodec:           "libx264",
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   1,
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)
		at.Equal(len(profile), 8)
		at.Equal(profile["-vcodec"], "libx264")
		at.Equal(profile["-r"], "30")
		at.Equal(profile["-force_key_frames"], "\"expr:gte(t,n_forced*10)\"")
		at.Equal(profile["-bufsize"], "100k")
		at.Equal(profile["-maxrate"], "100k")
		at.Equal(profile["-minrate"], "100k")
		at.Equal(profile["-crf"], "30")
		at.Equal(profile["-an"], "")
	}
	{
		codec_profile := CodecProfile{
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   2,
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		_, err := Transform(codec_profile, watermark)
		at.NotNil(err)
	}
	{
		codec_profile := CodecProfile{
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   2,
			Acodec:           "aac",
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)
		at.Equal(len(profile), 5)
		at.Equal(profile["-acodec"], "aac")
		at.Equal(profile["-vn"], "")
		at.Equal(profile["-b:a"], "128k")
		at.Equal(profile["-ac"], "1")
		at.Equal(profile["-ar"], "8000")
	}
	{
		codec_profile := CodecProfile{
			Vcodec:           "libx264",
			VideoLongSide:    250,
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   1,
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}
		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)
		at.Equal(len(profile), 9)
		at.Equal(profile["-vcodec"], "libx264")
		at.Equal(profile["-r"], "30")
		at.Equal(profile["-force_key_frames"], "\"expr:gte(t,n_forced*10)\"")
		at.Equal(profile["-bufsize"], "100k")
		at.Equal(profile["-maxrate"], "100k")
		at.Equal(profile["-minrate"], "100k")
		at.Equal(profile["-crf"], "30")
		at.Equal(profile["-an"], "")
		at.Equal(profile["-vf"], "\"scale='if(gt(250, max(iw\\, ih)), iw, 2*floor(250/max(iw\\, ih)*iw/2))':'if(gt(250, max(iw\\, ih)), ih, 2*floor(250/max(iw\\, ih)*ih/2))'\"")
	}
	{
		codec_profile := CodecProfile{
			Vcodec:           "libx264",
			VideoLongSide:    250,
			VideoBitRate:     200,
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			RemoveBFrame:     true,
			AudioVideoType:   0,
			Acodec:           "aac",
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{
			Image:           "water.png",
			ImageHigh:       320,
			ImageWidth:      240,
			X:               0.5,
			Y:               0.5,
			WidthProportion: 0.1,
		}

		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)

		at.Equal(len(profile), 14)
		at.Equal(profile["-vcodec"], "libx264")
		at.Equal(profile["-b:a"], "128k")
		at.Equal(profile["-force_key_frames"], "\"expr:gte(t,n_forced*10)\"")
		at.Equal(profile["-b:v"], "200k")
		at.Equal(profile["-acodec"], "aac")
		at.Equal(profile["-minrate"], "100k")
		at.Equal(profile["-bufsize"], "100k")
		at.Equal(profile["-vf"], "\"movie=water.png[wmo];[wmo][in]scale2ref=(main_w*0.100000):((320.000000/240.000000)*(main_w*0.100000))[wm][ino];[ino][wm]overlay=0.500000*main_w:0.500000*main_h[wmimage];[wmimage]scale='if(gt(250, max(iw\\, ih)), iw, 2*floor(250/max(iw\\, ih)*iw/2))':'if(gt(250, max(iw\\, ih)), ih, 2*floor(250/max(iw\\, ih)*ih/2))'[out]\"")
		at.Equal(profile["-maxrate"], "100k")
		at.Equal(profile["-crf"], "30")
		at.Equal(profile["-ac"], "1")
		at.Equal(profile["-x264opts"], "\"bframes=0\"")
		at.Equal(profile["-r"], "30")
		at.Equal(profile["-ar"], "8000")
	}
	{
		codec_profile := CodecProfile{
			Vcodec:           "libx264",
			VideoFrameRate:   30,
			KeyframeInterval: 10,
			AudioVideoType:   0,
			Acodec:           "aac",
			Ac:               1,
			Ab:               128,
			Ar:               8000,
			BufSize:          100,
			MaxRate:          100,
			MinRate:          100,
			Crf:              30,
		}

		watermark := WatermarkProfile{}

		profile, err := Transform(codec_profile, watermark)
		at.Nil(err)
		at.Equal(len(profile), 11)
		at.Equal(profile["-vcodec"], "libx264")
		at.Equal(profile["-b:a"], "128k")
		at.Equal(profile["-force_key_frames"], "\"expr:gte(t,n_forced*10)\"")
		at.Equal(profile["-acodec"], "aac")
		at.Equal(profile["-minrate"], "100k")
		at.Equal(profile["-bufsize"], "100k")
		at.Equal(profile["-maxrate"], "100k")
		at.Equal(profile["-crf"], "30")
		at.Equal(profile["-ac"], "1")
		at.Equal(profile["-r"], "30")
		at.Equal(profile["-ar"], "8000")
		at.Empty(profile["-vf"])
	}
}
