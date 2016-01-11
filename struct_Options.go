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

// Options
package main

type Options struct {
	Source      string
	Destination string

	JpegOutput         bool
	PngOutput          bool
	WebpOutputLossless bool
	WebpOutputLossy    bool

	JpegQuality         int
	WebpQualityLossless int
	WebpQualityLossy    int

	JpegDestination         string
	PngDestination          string
	WebpDestinationLossless string
	WebpDestinationLossy    string

	CjpegOptions         string
	CwebpOptionsLossless string
	CwebpOptionsLossy    string

	Folders bool
	MaxCpus int
}
