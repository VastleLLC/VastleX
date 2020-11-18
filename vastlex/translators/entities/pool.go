package entities

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

// Translator is an entity runtime id translater.
type Translater interface {
	Translate(pk packet.Packet, oldeid, neweid, olduid, newuid int64)
}

// Pool is the pool of translaters.
var Pool = map[uint32]func() Translater{
	packet.IDLogin:                      func() Translater { return &Login{} },
	packet.IDPlayStatus:                 func() Translater { return &PlayStatus{} },
	packet.IDServerToClientHandshake:    func() Translater { return &ServerToClientHandshake{} },
	packet.IDClientToServerHandshake:    func() Translater { return &ClientToServerHandshake{} },
	packet.IDDisconnect:                 func() Translater { return &Disconnect{} },
	packet.IDResourcePacksInfo:          func() Translater { return &ResourcePacksInfo{} },
	packet.IDResourcePackStack:          func() Translater { return &ResourcePackStack{} },
	packet.IDResourcePackClientResponse: func() Translater { return &ResourcePackClientResponse{} },
	packet.IDText:                       func() Translater { return &Text{} },
	packet.IDSetTime:                    func() Translater { return &SetTime{} },
	packet.IDStartGame:                  func() Translater { return &StartGame{} },
	packet.IDAddPlayer:                  func() Translater { return &AddPlayer{} },
	packet.IDAddActor:                   func() Translater { return &AddActor{} },
	packet.IDRemoveActor:                func() Translater { return &RemoveActor{} },
	packet.IDAddItemActor:               func() Translater { return &AddItemActor{} },
	// ---
	packet.IDTakeItemActor:     func() Translater { return &TakeItemActor{} },
	packet.IDMoveActorAbsolute: func() Translater { return &MoveActorAbsolute{} },
	packet.IDMovePlayer:        func() Translater { return &MovePlayer{} },
	packet.IDRiderJump:         func() Translater { return &RiderJump{} },
	packet.IDUpdateBlock:       func() Translater { return &UpdateBlock{} },
	packet.IDAddPainting:       func() Translater { return &AddPainting{} },
	packet.IDTickSync:          func() Translater { return &TickSync{} },
	// ---
	packet.IDLevelEvent:                  func() Translater { return &LevelEvent{} },
	packet.IDBlockEvent:                  func() Translater { return &BlockEvent{} },
	packet.IDActorEvent:                  func() Translater { return &ActorEvent{} },
	packet.IDMobEffect:                   func() Translater { return &MobEffect{} },
	packet.IDUpdateAttributes:            func() Translater { return &UpdateAttributes{} },
	packet.IDInventoryTransaction:        func() Translater { return &InventoryTransaction{} },
	packet.IDMobEquipment:                func() Translater { return &MobEquipment{} },
	packet.IDMobArmourEquipment:          func() Translater { return &MobArmourEquipment{} },
	packet.IDInteract:                    func() Translater { return &Interact{} },
	packet.IDBlockPickRequest:            func() Translater { return &BlockPickRequest{} },
	packet.IDActorPickRequest:            func() Translater { return &ActorPickRequest{} },
	packet.IDPlayerAction:                func() Translater { return &PlayerAction{} },
	packet.IDActorFall:                   func() Translater { return &ActorFall{} },
	packet.IDHurtArmour:                  func() Translater { return &HurtArmour{} },
	packet.IDSetActorData:                func() Translater { return &SetActorData{} },
	packet.IDSetActorMotion:              func() Translater { return &SetActorMotion{} },
	packet.IDSetActorLink:                func() Translater { return &SetActorLink{} },
	packet.IDSetHealth:                   func() Translater { return &SetHealth{} },
	packet.IDSetSpawnPosition:            func() Translater { return &SetSpawnPosition{} },
	packet.IDAnimate:                     func() Translater { return &Animate{} },
	packet.IDRespawn:                     func() Translater { return &Respawn{} },
	packet.IDContainerOpen:               func() Translater { return &ContainerOpen{} },
	packet.IDContainerClose:              func() Translater { return &ContainerClose{} },
	packet.IDPlayerHotBar:                func() Translater { return &PlayerHotBar{} },
	packet.IDInventoryContent:            func() Translater { return &InventoryContent{} },
	packet.IDInventorySlot:               func() Translater { return &InventorySlot{} },
	packet.IDContainerSetData:            func() Translater { return &ContainerSetData{} },
	packet.IDCraftingData:                func() Translater { return &CraftingData{} },
	packet.IDCraftingEvent:               func() Translater { return &CraftingEvent{} },
	packet.IDGUIDataPickItem:             func() Translater { return &GUIDataPickItem{} },
	packet.IDAdventureSettings:           func() Translater { return &AdventureSettings{} },
	packet.IDBlockActorData:              func() Translater { return &BlockActorData{} },
	packet.IDPlayerInput:                 func() Translater { return &PlayerInput{} },
	packet.IDLevelChunk:                  func() Translater { return &LevelChunk{} },
	packet.IDSetCommandsEnabled:          func() Translater { return &SetCommandsEnabled{} },
	packet.IDSetDifficulty:               func() Translater { return &SetDifficulty{} },
	packet.IDChangeDimension:             func() Translater { return &ChangeDimension{} },
	packet.IDSetPlayerGameType:           func() Translater { return &SetPlayerGameType{} },
	packet.IDPlayerList:                  func() Translater { return &PlayerList{} },
	packet.IDSimpleEvent:                 func() Translater { return &SimpleEvent{} },
	packet.IDEvent:                       func() Translater { return &Event{} },
	packet.IDSpawnExperienceOrb:          func() Translater { return &SpawnExperienceOrb{} },
	packet.IDClientBoundMapItemData:      func() Translater { return &ClientBoundMapItemData{} },
	packet.IDMapInfoRequest:              func() Translater { return &MapInfoRequest{} },
	packet.IDRequestChunkRadius:          func() Translater { return &RequestChunkRadius{} },
	packet.IDChunkRadiusUpdated:          func() Translater { return &ChunkRadiusUpdated{} },
	packet.IDItemFrameDropItem:           func() Translater { return &ItemFrameDropItem{} },
	packet.IDGameRulesChanged:            func() Translater { return &GameRulesChanged{} },
	packet.IDCamera:                      func() Translater { return &Camera{} },
	packet.IDBossEvent:                   func() Translater { return &BossEvent{} },
	packet.IDShowCredits:                 func() Translater { return &ShowCredits{} },
	packet.IDAvailableCommands:           func() Translater { return &AvailableCommands{} },
	packet.IDCommandRequest:              func() Translater { return &CommandRequest{} },
	packet.IDCommandBlockUpdate:          func() Translater { return &CommandBlockUpdate{} },
	packet.IDCommandOutput:               func() Translater { return &CommandOutput{} },
	packet.IDUpdateTrade:                 func() Translater { return &UpdateTrade{} },
	packet.IDUpdateEquip:                 func() Translater { return &UpdateEquip{} },
	packet.IDResourcePackDataInfo:        func() Translater { return &ResourcePackDataInfo{} },
	packet.IDResourcePackChunkData:       func() Translater { return &ResourcePackChunkData{} },
	packet.IDResourcePackChunkRequest:    func() Translater { return &ResourcePackChunkRequest{} },
	packet.IDTransfer:                    func() Translater { return &Transfer{} },
	packet.IDPlaySound:                   func() Translater { return &PlaySound{} },
	packet.IDStopSound:                   func() Translater { return &StopSound{} },
	packet.IDSetTitle:                    func() Translater { return &SetTitle{} },
	packet.IDAddBehaviourTree:            func() Translater { return &AddBehaviourTree{} },
	packet.IDStructureBlockUpdate:        func() Translater { return &StructureBlockUpdate{} },
	packet.IDShowStoreOffer:              func() Translater { return &ShowStoreOffer{} },
	packet.IDPurchaseReceipt:             func() Translater { return &PurchaseReceipt{} },
	packet.IDPlayerSkin:                  func() Translater { return &PlayerSkin{} },
	packet.IDSubClientLogin:              func() Translater { return &SubClientLogin{} },
	packet.IDAutomationClientConnect:     func() Translater { return &AutomationClientConnect{} },
	packet.IDSetLastHurtBy:               func() Translater { return &SetLastHurtBy{} },
	packet.IDBookEdit:                    func() Translater { return &BookEdit{} },
	packet.IDNPCRequest:                  func() Translater { return &NPCRequest{} },
	packet.IDPhotoTransfer:               func() Translater { return &PhotoTransfer{} },
	packet.IDModalFormRequest:            func() Translater { return &ModalFormRequest{} },
	packet.IDModalFormResponse:           func() Translater { return &ModalFormResponse{} },
	packet.IDServerSettingsRequest:       func() Translater { return &ServerSettingsRequest{} },
	packet.IDServerSettingsResponse:      func() Translater { return &ServerSettingsResponse{} },
	packet.IDShowProfile:                 func() Translater { return &ShowProfile{} },
	packet.IDSetDefaultGameType:          func() Translater { return &SetDefaultGameType{} },
	packet.IDRemoveObjective:             func() Translater { return &RemoveObjective{} },
	packet.IDSetDisplayObjective:         func() Translater { return &SetDisplayObjective{} },
	packet.IDSetScore:                    func() Translater { return &SetScore{} },
	packet.IDLabTable:                    func() Translater { return &LabTable{} },
	packet.IDUpdateBlockSynced:           func() Translater { return &UpdateBlockSynced{} },
	packet.IDMoveActorDelta:              func() Translater { return &MoveActorDelta{} },
	packet.IDSetScoreboardIdentity:       func() Translater { return &SetScoreboardIdentity{} },
	packet.IDSetLocalPlayerAsInitialised: func() Translater { return &SetLocalPlayerAsInitialised{} },
	packet.IDUpdateSoftEnum:              func() Translater { return &UpdateSoftEnum{} },
	packet.IDNetworkStackLatency:         func() Translater { return &NetworkStackLatency{} },
	// ---
	packet.IDScriptCustomEvent:           func() Translater { return &ScriptCustomEvent{} },
	packet.IDSpawnParticleEffect:         func() Translater { return &SpawnParticleEffect{} },
	packet.IDAvailableActorIdentifiers:   func() Translater { return &AvailableActorIdentifiers{} },
	packet.IDNetworkChunkPublisherUpdate: func() Translater { return &NetworkChunkPublisherUpdate{} },
	packet.IDBiomeDefinitionList:         func() Translater { return &BiomeDefinitionList{} },
	packet.IDLevelSoundEvent:             func() Translater { return &LevelSoundEvent{} },
	packet.IDLevelEventGeneric:           func() Translater { return &LevelEventGeneric{} },
	packet.IDLecternUpdate:               func() Translater { return &LecternUpdate{} },
	// ---
	packet.IDAddEntity:                         func() Translater { return &AddEntity{} },
	packet.IDRemoveEntity:                      func() Translater { return &RemoveEntity{} },
	packet.IDClientCacheStatus:                 func() Translater { return &ClientCacheStatus{} },
	packet.IDOnScreenTextureAnimation:          func() Translater { return &OnScreenTextureAnimation{} },
	packet.IDMapCreateLockedCopy:               func() Translater { return &MapCreateLockedCopy{} },
	packet.IDStructureTemplateDataRequest:      func() Translater { return &StructureTemplateDataResponse{} },
	packet.IDStructureTemplateDataResponse:     func() Translater { return &StructureTemplateDataResponse{} },
	packet.IDUpdateBlockProperties:             func() Translater { return &UpdateBlockProperties{} },
	packet.IDClientCacheBlobStatus:             func() Translater { return &ClientCacheBlobStatus{} },
	packet.IDClientCacheMissResponse:           func() Translater { return &ClientCacheMissResponse{} },
	packet.IDEducationSettings:                 func() Translater { return &EducationSettings{} },
	packet.IDEmote:                             func() Translater { return &Emote{} },
	packet.IDMultiPlayerSettings:               func() Translater { return &MultiPlayerSettings{} },
	packet.IDSettingsCommand:                   func() Translater { return &SettingsCommand{} },
	packet.IDAnvilDamage:                       func() Translater { return &AnvilDamage{} },
	packet.IDCompletedUsingItem:                func() Translater { return &CompletedUsingItem{} },
	packet.IDNetworkSettings:                   func() Translater { return &NetworkSettings{} },
	packet.IDPlayerAuthInput:                   func() Translater { return &PlayerAuthInput{} },
	packet.IDCreativeContent:                   func() Translater { return &CreativeContent{} },
	packet.IDPlayerEnchantOptions:              func() Translater { return &PlayerEnchantOptions{} },
	packet.IDItemStackRequest:                  func() Translater { return &ItemStackRequest{} },
	packet.IDItemStackResponse:                 func() Translater { return &ItemStackResponse{} },
	packet.IDPlayerArmourDamage:                func() Translater { return &PlayerArmourDamage{} },
	packet.IDCodeBuilder:                       func() Translater { return &CodeBuilder{} },
	packet.IDUpdatePlayerGameType:              func() Translater { return &UpdatePlayerGameType{} },
	packet.IDEmoteList:                         func() Translater { return &EmoteList{} },
	packet.IDPositionTrackingDBServerBroadcast: func() Translater { return &PositionTrackingDBServerBroadcast{} },
	packet.IDPositionTrackingDBClientRequest:   func() Translater { return &PositionTrackingDBClientRequest{} },
	packet.IDDebugInfo:                         func() Translater { return &DebugInfo{} },
	packet.IDPacketViolationWarning:            func() Translater { return &PacketViolationWarning{} },
}
