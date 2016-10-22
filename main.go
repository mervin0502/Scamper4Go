package main

import (
	// "log"
	"mervin.me/Scamper4Go/extract"
)

func main() {
	// extract.AnalysisDump("./data/test.warts.creating")
	// extract.AnalysisDump("./data/test2.warts.creating")
	// extract.AnalysisDump("./data/test3.warts.creating")
	// extract.AnalysisDump("./data/test4.warts.creating")
	// extract.AnalysisDump("./data/test.txt", "./data/test.warts.creating")
	// extract.AnalysisDump("./data/test.txt", "./data/test3.gz")
	// extract.AnalysisDump("./data/test00.txt", "./data/test00.warts.creating")
	extract.TopologyDump("./data/test00.txt", "./data/test00.warts.creating")
	// extract.AnalysisDump("./data/test2.txt", "./data/test2.warts.creating")
	// extract.AnalysisDump("./data/test3.gz")
}
