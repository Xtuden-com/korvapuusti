/*
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 */
package main

import (
	"log"

	"flag"

	"github.com/google-research/korvapuusti/analysis/data_processing/tfexamples"
)

var (
	filePattern     = flag.String("filepattern", "ood_white", "The file pattern of .wav files to read.")
	dataPath        = flag.String("datapath", "ood_white_white/preprocessed_ood_test_set.json", "The file with the data to prepare examples for.")
	fftWindowSize   = flag.Int("fftWindowSize", 2048, "Window size of FFT to use.")
	carfacStableIdx = flag.Int("carfacStableIdx", 2205, "How many featuredimensions to skip from the CARFAC features, assuming the calculations will be stable after.")
	skipSeconds     = flag.Float64("skipSeconds", 1., "How many seconds of the .wav files to skip.")
	sampleRate      = flag.Int("sampleRate", 48000, "The sample rate in the .wav files.")
	shardSize       = flag.Int("shardSize", 20, "How many examples to put in each shard.")
	filePrefix      = flag.String("filePrefix", "tfdata/oodwhiteraw", "With what prefix to save the output shards.")
	unityDBLevel    = flag.Float64("unityDBLevel", 100.0, "Level of a unity sine.")
	extractMetadata = flag.Bool("extractmetadata", false, "Whether the wavfilepaths contain meta data on the specifics of the contained sounds.")
)

func main() {
	err := tfexamples.WriteExamplesToRecordIO(*filePattern, *dataPath, *filePrefix, *shardSize, *carfacStableIdx,
		*fftWindowSize, *sampleRate, *skipSeconds, *unityDBLevel, *extractMetadata)
	if err != nil {
		log.Panicf("Function WriteExamplesToRecordIO failed with error:  %+v", err)
	}
}
