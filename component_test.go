package es

import (
	"testing"

	"github.com/futile/go-lil-t"
)

type mockComponent struct {
	BaseComponent

	mockBool bool
}

func mockComponentFactory() Component {
	return &mockComponent{mockBool: false}
}

func TestComponentContainer(t *testing.T) {
	If := lilt.NewIf(t)

	e := Entity{id: 0, reuseCount: 0}

	cc := NewComponentContainer(mockComponentFactory)
	If(cc.Has(e)).Errorf("Has() returned true for empty ComponentContainer!")

	c, err := cc.Create(e)
	If(c == nil).Errorf("Create() returned nil!")
	If(err != nil).Errorf("Create() returned an error!")
	If(!cc.Has(e)).Errorf("Has() returned false even though Component was created!")

	if _, ok := c.(*mockComponent); !ok {
		t.Errorf("Create() returned wrong type! expected: *mockComponent, but got: %T", c)
	}

	e.reuseCount = 1
	If(cc.Has(e)).Errorf("Has() returned true even though reuseCount was changed!")

	c, _ = cc.Create(e)
	If(c.reuseCount() != e.reuseCount).Errorf("Wrong reuse counts: entity: %v, component: %v", e.reuseCount, c.reuseCount())

	_, err = cc.Create(e)
	If(err == nil).Errorf("Second call to Create() did not cause an error!")

	c2 := cc.Get(e)
	If(c2 != c).Errorf("Get() returned a different Component than Create() did!")

	c2 = cc.GetOrCreate(e)
	If(c2 != c).Errorf("GetOrCreate() returned a different Component than Create() did!")

	err = cc.Remove(e)
	If(err != nil).Errorf("Remove() falsely returned an error: %v", err)
	If(cc.Has(e)).Errorf("Has() returned true after call to Remove()!")

	err = cc.Remove(e)
	If(err == nil).Errorf("Second call to Remove() did not return an error!")
}
