package main

import (
	_ "github.com/corona10/goimagehash"
	"github.com/giorgisio/goav/avformat"
	_ "github.com/multiformats/go-multihash"
)

func init() {
	avformat.AvRegisterAll()
}

func vphash_file(videopath string) {
	pFormatContext := avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&pFormatContext, videopath, nil, nil) != 0 {
		Log.FatalF("Unable to open file: %s\n", videopath)
	}
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		Log.FatalF("Couldn't find stream information")
	}
	pFormatContext.AvDumpFormat(0, videopath, 0)

	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		switch pFormatContext.Streams()[i].CodecParameters().AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:
			// got video stream
			vphash_stream(pFormatContext.Streams()[i])
			return
		default:
			continue
		}
	}
	Log.FatalF("Couldn't find any video stream")
}

func vphash_stream(stream *avformat.Stream) {
	pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(stream.GetCodecId()))
			if pCodec == nil {
				Log.FatalF("Unsupported codec!")
			}
			// Copy context
			pCodecCtx := pCodec.AvcodecAllocContext3()
			if pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(stream))) != 0 {
				Log.FatalF("Couldn't copy codec context")
			}

			// Open codec
			if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
				Log.FatalF("Could not open codec")
			}

			// Allocate video frame
			pFrame := avutil.AvFrameAlloc()

			// Allocate an AVFrame structure
			pFrameRGB := avutil.AvFrameAlloc()
			if pFrameRGB == nil {
				Log.FatalF("Unable to allocate RGB Frame")
			}

}
