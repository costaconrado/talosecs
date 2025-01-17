package talosecs

import "reflect"

var isInitialized bool

// Init initializes all of your ECS updateSystems. Call it once on game world start before calling Update
func Init() {
	if isInitialized {
		panic("ECS already initialized!")
	}

	initLayers()
	isInitialized = true
}

// Update calls Update of each UpdateSystem every frame. Used to handle most of game logic.
func Update() {
	updateLayers()
	clearOneFrames()
}

// Reset returns state of whole ECS to default - kills all entities, clears all layers with systems. Can be used for tests and some specific purposes.
func Reset() {
	ResetEcsState()
	layers = nil

	currentEntityId = 0
	isInitialized = false
}

// ResetEcsState removes all entities, components and signals. Can be called to restart game scene etc
func ResetEcsState() {
	componentsPool = map[reflect.Type][]any{}
	entsComponents = map[Entity][]any{}
	componentsEnts = map[any]Entity{}
	signals = nil
	oneFrames = nil
}
