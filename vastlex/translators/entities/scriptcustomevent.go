package entities

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type ScriptCustomEvent struct {}

func (ScriptCustomEvent) Translate(pk packet.Packet, eid1, eid2 int64, uid1, uid2 int64) {
}