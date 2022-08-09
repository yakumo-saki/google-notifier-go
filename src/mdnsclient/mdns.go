package mdnsclient

import (
	"container/list"
	"context"
	"fmt"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/yakumo-saki/google-notifier-go/src/config"
)

func Scan() {
	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		l.Err(err).Msg("Failed to initialize resolver")
	}

	entries := list.New()
	entriesChan := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry, list *list.List) {
		for entry := range results {
			l.Debug().Msgf("%s", entry)
			list.PushBack(entry)
		}
	}(entriesChan, entries)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = resolver.Browse(ctx, "_googlecast._tcp", "local.", entriesChan)
	if err != nil {
		l.Err(err).Msg("Failed to browse")
	}

	l.Info().Msgf("Wait for mDNS reply...")
	<-ctx.Done()

	for e := entries.Front(); e != nil; e = e.Next() {
		ent := e.Value.(*zeroconf.ServiceEntry)
		fmt.Printf("Instance: %s", ent.Instance)
		fmt.Printf(" Hostname: %s", ent.HostName)
		fmt.Printf(" IP4: [")
		for i, ip := range ent.AddrIPv4 {
			fmt.Printf("(%d) %s", i, ip)
		}
		fmt.Printf("]")

		if config.IsIgnoredInstance(ent.Instance) {
			fmt.Printf(" *Ignored*")
		}

		fmt.Printf("\n")
	}

}
