package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestText_Selected(t *testing.T) {
	e := widget.NewEntry()
	test.Type(e, "Hello")
	assert.Equal(t, "Hello", e.Text)

	test.DoubleTap(e)
	assert.Equal(t, "Hello", e.SelectedText())
	assert.Equal(t, 5, e.CursorColumn)
}
