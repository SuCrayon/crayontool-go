package main

import (
	"crayontool-go/pkg/strutil"
	"log"
)

const (
	bannerTemplate = "%s" +
		"                                     __              __                %s" +
		"  ______________ ___  ______  ____  / /_____  ____  / /     ____ _____ %s" +
		" / ___/ ___/ __ `/ / / / __ \\/ __ \\/ __/ __ \\/ __ \\/ /_____/ __ `/ __ \\%s" +
		"/ /__/ /  / /_/ / /_/ / /_/ / / / / /_/ /_/ / /_/ / /_____/ /_/ / /_/ /%s" +
		"\\___/_/   \\__,_/\\__, /\\____/_/ /_/\\__/\\____/\\____/_/      \\__, /\\____/ %s" +
		"               /____/                                    /____/        %s"
)

func printBanner() {
	log.Print(strutil.SprintfRepeatedTimes(bannerTemplate, strutil.GetLineSep(), 7))
}

func init() {
	printBanner()
}
