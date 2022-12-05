package mdnsclient

import (
	"container/list"
	"context"
	"fmt"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/yakumo-saki/google-notifier-go/src/config"
)

func Scan(withPrint bool) []*zeroconf.ServiceEntry {
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

	services := make([]*zeroconf.ServiceEntry, 0, entries.Len())
	for e := entries.Front(); e != nil; e = e.Next() {
		ent := e.Value.(*zeroconf.ServiceEntry)
		if withPrint {
			printServiceEntry(ent)
		}

		if config.IsIgnoredInstance(ent.Instance) {
			continue
		}

		services = append(services, ent)
	}

	return services
}

func printServiceEntry(service *zeroconf.ServiceEntry) {
	fmt.Printf("Instance: %s", service.Instance)
	fmt.Printf(" Hostname: %s", service.HostName)
	fmt.Printf(" IP4: [")
	for i, ip := range service.AddrIPv4 {
		fmt.Printf("(%d) %s", i, ip)
	}
	fmt.Printf("]")

	if config.IsIgnoredInstance(service.Instance) {
		fmt.Printf(" *Ignored*")
	}

	fmt.Printf("\n")
}
