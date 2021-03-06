package net

import (
	"context"
	"fmt"
	"testing"

	p2pty "github.com/33cn/chain33/system/p2p/dht/types"
)

func Test_initInnerPeers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hosts := getNetHosts(ctx, 8, t)

	h0 := hosts[0].Peerstore().PeerInfo(hosts[0].ID())
	h1 := hosts[1].Peerstore().PeerInfo(hosts[1].ID())
	h2 := hosts[2].Peerstore().PeerInfo(hosts[2].ID())
	h3 := hosts[3].Peerstore().PeerInfo(hosts[3].ID())
	h4 := hosts[4].Peerstore().PeerInfo(hosts[4].ID())
	h7 := hosts[7].Peerstore().PeerInfo(hosts[7].ID())
	h0str := h0.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h0.ID.Pretty())
	h1str := h1.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h1.ID.Pretty())
	h2str := h2.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h2.ID.Pretty())
	h3str := h3.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h3.ID.Pretty())
	h4str := h4.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h4.ID.Pretty())
	h7str := h7.Addrs[0].String() + fmt.Sprintf("/p2p/%v", h7.ID.Pretty())
	subcfg := &p2pty.P2PSubConfig{}
	subcfg.BootStraps = []string{h1str, h2str, h3str, h4str}
	subcfg.Seeds = []string{h0str, h7str}

	initInnerPeers(hosts[5], nil, subcfg)
	hosts[5].Close()

}
