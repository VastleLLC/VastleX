package vastlex

import (
	"errors"
	"fmt"
	"github.com/VastleLLC/VastleX/config"
	"github.com/VastleLLC/VastleX/log"
	"github.com/VastleLLC/VastleX/vastlex/server"
	"github.com/VastleLLC/VastleX/vastlex/session"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"net/http"
	_ "net/http/pprof"
)

// VastleX is the main structure for the proxy.
var VastleX Proxy = vastlex

var vastlex = &Structure{
	listener: &minecraft.Listener{
		ErrorLog:               nil,
		AuthenticationDisabled: !config.Config.Minecraft.Auth,
		ServerName:             config.Config.Minecraft.Motd,
		ShowVersion:            config.Config.Minecraft.ShowVersion,
	},
	info: server.Info{
		Host: config.Config.Listener.Host,
		Port: config.Config.Listener.Port,
	},
	players: map[string]Player{},
}

// Structure is the structure of VastleX.
type Structure struct {
	listener *minecraft.Listener
	info     server.Info
	players  map[string]Player
}

// Start starts the proxy.
func Start() (err error) {
	if config.Config.Debug.Profiling {
		go func() {
			log.FatalError("Error starting profiling server", http.ListenAndServe(config.Config.Listener.Host + ":6060", nil))
		}()
	}
	err = vastlex.listener.Listen("raknet", fmt.Sprintf("%v:%v", vastlex.info.Host, vastlex.info.Port))
	if err != nil {
		return err
	}
	log.Info().Str("host", vastlex.info.Host).Int("port", vastlex.info.Port).Str("checksum", log.Checksum).Msg("VastleX is listening for players")

	for {
		conn, err := vastlex.listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			vastlex.players[conn.(*minecraft.Conn).IdentityData().DisplayName] = session.New(conn.(*minecraft.Conn))
			log.Info().Str("username", conn.(*minecraft.Conn).IdentityData().DisplayName).Msg("Player connected")
			if config.Config.Lobby.Enabled {
				err = vastlex.players[conn.(*minecraft.Conn).IdentityData().DisplayName].Send(server.Info{
					Host: config.Config.Lobby.Host,
					Port: config.Config.Lobby.Port,
				})
				if err != nil {
					vastlex.players[conn.(*minecraft.Conn).IdentityData().DisplayName].Kick(text.Red()("We had an error connecting you to a lobby"))
					log.Err().Str("username", conn.(*minecraft.Conn).IdentityData().DisplayName).Err(err).Msg("Player failed to connect to lobby")
				}
			}
		}()
	}
}

// ...
func (vastlex *Structure) Config() config.Structure {
	return config.Config
}

// ...
func (vastlex *Structure) Motd() string {
	return vastlex.listener.ServerName
}

// ...
func (vastlex *Structure) SetMotd(motd string) {
	vastlex.listener.ServerName = motd
}

// ...
func (vastlex *Structure) Players() map[string]Player {
	return vastlex.players
}

// ...
func (vastlex *Structure) GetPlayer(username string) (Player, error) {
	if vastlex.players[username] != nil {
		return vastlex.players[username], nil
	} else {
		return nil, errors.New("player not found")
	}
}
