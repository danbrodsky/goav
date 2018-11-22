package main

import (
	"log"

	"github.com/danbrodsky/goav/avcodec"
	"github.com/danbrodsky/goav/avdevice"
	"github.com/danbrodsky/goav/avfilter"
	"github.com/danbrodsky/goav/avformat"
	"github.com/danbrodsky/goav/avutil"
	"github.com/danbrodsky/goav/swresample"
	"github.com/danbrodsky/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
