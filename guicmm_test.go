﻿package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_makeUI(t *testing.T) {
	var testCfg config
	edit, preview := testCfg.makeUI()
	test.Type(edit, "Hello")
	if preview.String() != "Hello" {
		t.Error("failed -- did not find expected value in preview")
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test markdown")
	edit, preview := testCfg.makeUI()
	testCfg.createMenuItems(testWin)
	testWin.SetContent(container.NewVSplit(edit, preview))
	testApp.Run()
	test.Type(edit, "some test")
	if preview.String() != "some test" {
		t.Error("Test failed")
	}
}
