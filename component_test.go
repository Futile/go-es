package es

import "testing"

type mockComponent struct {
	BaseComponent

	mockBool bool
}

var mockComponentFactory ComponentFactory = func() Component {
	return &mockComponent{mockBool: false}
}

func TestComponentContainer(t *testing.T) {
	w := NewWorld()
	e := w.NewEntity()

	cc := NewComponentContainer(mockComponentFactory)

	if cc.Has(e) {
		t.Errorf("Has() returned true for empty ComponentContainer!")
	}

	c, err := cc.Create(e)

	if c == nil {
		t.Errorf("Create() returned nil!")
	}

	if err != nil {
		t.Errorf("Create() returned an error!")
	}

	if _, ok := c.(*mockComponent); !ok {
		t.Errorf("Create() returned wrong type! expected: *mockComponent, but got: %T", c)
	}

	if !cc.Has(e) {
		t.Errorf("Has() returned false even though Component was created!")
	}

	w.RemoveEntity(e)

	e = w.NewEntity()

	if cc.Has(e) {
		t.Errorf("Has() returned true even though entity was re-created!")
	}

	c, _ = cc.Create(e)

	if c.reuseCount() != e.reuseCount {
		t.Errorf("Wrong reuse counts: entity: %v, component: %v", e.reuseCount, c.reuseCount())
	}

	_, err = cc.Create(e)

	if err == nil {
		t.Errorf("Second call to Create() did not cause an error!")
	}

	c2 := cc.Get(e)

	if c2 != c {
		t.Errorf("Get() returned a different Component than Create() did!")
	}

	c2 = cc.GetOrCreate(e)

	if c2 != c {
		t.Errorf("GetOrCreate() returned a different Component than Create() did!")
	}

	err = cc.Remove(e)

	if err != nil {
		t.Errorf("Remove() falsely returned an error: %v", err)
	}

	if cc.Has(e) {
		t.Errorf("Has() returned true after call to Remove()!")
	}

	err = cc.Remove(e)

	if err == nil {
		t.Errorf("Second call to Remove() did not return an error!")
	}
}
