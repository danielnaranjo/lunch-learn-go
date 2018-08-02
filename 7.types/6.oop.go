package main

import "fmt"

type Positionable struct {
	X uint8
	Y uint8
}

// A method on Positionable
func (positionable *Positionable) Debug() string {
	return fmt.Sprintf("<Positioned at (%v, %v)>", positionable.X, positionable.Y)
}

type Sized struct {
	Width uint8
	Height uint8
}

// A method on Sized
func (sized *Sized) GetArea() uint16 {
	return uint16(sized.Width) * uint16(sized.Height)
}

func (sized *Sized) Debug() string {
	return fmt.Sprintf("<Sized with (%v, %v)>", sized.Width, sized.Height)
}

type WindRoseObject struct {
	Positionable
	Sized
}

func (wrObject *WindRoseObject) Debug() string {
	return fmt.Sprintf("<WindRose %v,%v,%v,%v from %v and %v",
		               wrObject.X, wrObject.Y, wrObject.Width, wrObject.Height,
		               wrObject.Positionable.Debug(), wrObject.Sized.Debug())
}

type Debuggable interface {
	Debug() string
}

func DumpDebug(dbg Debuggable) {
	fmt.Println("*********************")
	fmt.Println(">> " + dbg.Debug())
	fmt.Println("*********************")
}

func main() {
	pos := Positionable{3, 4}
	sized := Sized{2, 2}
	DumpDebug(&pos)
	DumpDebug(&sized)
	wrObj := WindRoseObject{pos, sized}
	DumpDebug(&wrObj)
	// Example of assigning bound methods (remember: the syntactic sugar allows us to not &pointerize our wrObject)
	areaMethod := wrObj.GetArea
	fmt.Printf("Area: %v\n", areaMethod())
	// Example of assigning unbound methods (however: the syntactic sugar does not exist in the unbound case)
	anotherAreaMethod := (*WindRoseObject).GetArea
	fmt.Printf("Area: %v\n", anotherAreaMethod(&wrObj))
}