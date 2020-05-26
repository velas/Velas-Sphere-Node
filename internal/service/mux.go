package service

import (
	"log"
	"sync"
)

func Mux(wg *sync.WaitGroup, services map[string]Service) Service {
	return func() error {
		for label, service := range services {
			if service == nil {
				continue
			}

			log.Println("starting service \"" + label + "\"")
			wg.Add(1)
			go func(wg *sync.WaitGroup, label string, service Service) {
				defer func() {
					r := recover()
					if r != nil {
						log.Println("got panic:", r)
					}

					wg.Done()
				}()

				err := service()
				if err != nil {
					log.Println("service \""+label+"\" failed:", err)
				}
			}(wg, label, service)
		}

		return nil
	}
}
