/*
   Copyright 2016 Jonathan Windle

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// rawsync
package main

import (
	"flag"
)

func main() {
	options := new(Options)
	flag.StringVar(&options.Source,
		"s", "", "sets the source file or directory")
	flag.StringVar(&options.Destination,
		"d", "", "sets the destination file or directory")

	flag.BoolVar(&options.JpegOutput,
		"jpeg", false, "turns on JPEG output")
	flag.BoolVar(&options.PngOutput,
		"png", false, "turns on PNG output")
	flag.BoolVar(&options.WebpOutputLossless,
		"webp", false, "turns on WebP lossless output")
	flag.BoolVar(&options.WebpOutputLossy,
		"webpl", false, "turns on WebP lossy output")

	flag.IntVar(&options.JpegQuality,
		"qj", 90, "sets JPEG quality")
	flag.IntVar(&options.WebpQualityLossless,
		"qw", 100, "sets WebP lossless quality")
	flag.IntVar(&options.WebpQualityLossy,
		"qwl", 90, "sets WebP lossy quality")

	flag.StringVar(&options.JpegDestination,
		"dj", "", "sets the destination file or directory for JPEGs")
	flag.StringVar(&options.PngDestination,
		"dp", "", "sets the destination file or directory for PNGs")
	flag.StringVar(&options.WebpDestinationLossless,
		"dw", "", "sets the destination file or directory for Lossless WebPs")
	flag.StringVar(&options.WebpDestinationLossy,
		"dwl", "", "sets the destination file or directory for Lossy WebPs")

	flag.StringVar(&options.CjpegOptions,
		"optsj", "", "sets command line options for cjpeg")
	flag.StringVar(&options.CwebpOptionsLossless,
		"optsw", "-preset photo",
		"sets command line options for cwebp in lossless mode")
	flag.StringVar(&options.CwebpOptionsLossy,
		"optswl", "-preset photo",
		"sets command line options for cwebp in lossy mode")

	flag.BoolVar(&options.Folders,
		"folders", false, "enable destination folders for batch transfers")
	flag.IntVar(&options.MaxCpus,
		"maxcpus", 0, "sets the maximum cpus used to convert media")

	flag.Parse()
}
