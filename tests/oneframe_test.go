package talosEcsTests

import (
	"testing"

	"github.com/costaconrado/talosecs"
)

func TestOneFrameAdd(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	e.OneFrame(&CompA{})

	if _, ok := talosecs.GetComponent[*CompA](e); !ok {
		t.Fatalf(`No OneFrame component was added`)
	}
}

func TestOneFrameRemove(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	e.OneFrame(&CompA{})
	talosecs.Update()

	if _, ok := talosecs.GetComponent[*CompA](e); ok {
		t.Fatalf(`OneFrame was not removed after one frame was passed`)
	}
}
