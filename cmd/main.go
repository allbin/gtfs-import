package main

import (
	"github.com/allbin/gtfs2postgis/health"
	"github.com/allbin/gtfs2postgis/importer"
	"github.com/jasonlvhit/gocron"
	"log"
)

// healthAddr is probed by the deployment's readiness and liveness probes
// (k8s/deployment.yaml). The listener starts before the first import so a
// long import does not look like a dead container.
const healthAddr = ":8080"

func main() {
	go func() {
		log.Fatalf("health server stopped: %v", health.ListenAndServe(healthAddr))
	}()
	importer.Run()
	gocron.Every(1).Day().At("23:00").Do(importer.Run)
	_, t := gocron.NextRun()
	gocron.Start()
	log.Printf("Next run scheduled at %s.", t.Format("15:04:05"))
	select {}
}
