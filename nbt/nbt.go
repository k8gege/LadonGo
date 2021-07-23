package nbt
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/time/rate"
	"os"
	"runtime"
	"sync"
	"time"
)

type ScanResult struct {
	Host  string            `json:"host"`
	Port  string            `json:"port,omitempty"`
	Proto string            `json:"proto,omitempty"`
	Probe string            `json:"probe,omitempty"`
	Name  string            `json:"name,omitempty"`
	Nets  []string          `json:"nets,omitempty"`
	Info  map[string]string `json:"info"`
}

type Prober interface {
	Setup()
	Initialize()
	Wait()
	AddTarget(string)
	CloseInput()
	SetOutput(chan<- ScanResult)
	CheckRateLimit()
	SetLimiter(*rate.Limiter)
}

type Probe struct {
	name    string
	options map[string]string
	waiter  sync.WaitGroup
	input   chan string
	output  chan<- ScanResult
	limiter *rate.Limiter
}

func (this *Probe) String() string {
	return fmt.Sprintf("%s", this.name)
}

func (this *Probe) Wait() {
	this.waiter.Wait()
	return
}

func (this *Probe) Setup() {
	this.name = "generic"
	this.input = make(chan string)
	return
}

func (this *Probe) Initialize() {
	this.Setup()
	this.name = "generic"
	return
}

func (this *Probe) SetOutput(c_out chan<- ScanResult) {
	this.output = c_out
	return
}

func (this *Probe) AddTarget(t string) {
	this.input <- t
	return
}

func (this *Probe) CloseInput() {
	close(this.input)
	return
}

func (this *Probe) SetLimiter(limiter *rate.Limiter) {
	this.limiter = limiter
	return
}

func (this *Probe) CheckRateLimit() {
	for this.limiter.Allow() == false {
		time.Sleep(10 * time.Millisecond)
	}
}

var limiter *rate.Limiter
var ppsrate *int
var probes []Prober
var wi sync.WaitGroup
var wo sync.WaitGroup

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " [cidr] ... [cidr]")
	fmt.Println("")
	fmt.Println("Probes a list of networks for potential pivot points.")
	fmt.Println("")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func outputWriter(o <-chan ScanResult) {

	for found := range o {
		j, err := json.Marshal(found)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling result: '%v' : %s\n", found, err)
			continue
		}
		os.Stdout.Write(j)
		os.Stdout.Write([]byte("\n"))
	}
	wo.Done()
}

func initializeProbes(c_out chan<- ScanResult) {
	for _, probe := range probes {
		probe.Initialize()
		probe.SetOutput(c_out)
		probe.SetLimiter(limiter)
	}
}

func waitProbes() {
	for _, probe := range probes {
		probe.Wait()
	}
}

func processAddress(i <-chan string, o chan<- ScanResult) {
	for addr := range i {
		for _, probe := range probes {
			probe.AddTarget(addr)
		}
	}

	for _, probe := range probes {
		probe.CloseInput()
	}
	wi.Done()
}

func Info(target string) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	 flag.Usage = func() { usage() }
	 version := flag.Bool("version", false, "Show the application version")
	 ppsrate = flag.Int("rate", 1000, "Set the maximum packets per second rate")

	 flag.Parse()

if *version {
//printversion("nextnet")
 //os.exit(0)
}

	limiter = rate.NewLimiter(rate.Limit(*ppsrate), *ppsrate*3)

	// Input addresses
	c_addr := make(chan string)
	//fmt.Printf("type:%T value:%#v\n", c_addr, c_addr)



	// Output structs
	c_out := make(chan ScanResult)

	// Configure the probes
	initializeProbes(c_out)

			// for _, probe := range probes {
			// probe.AddTarget(target)
		// }
		
	// Launch a single input address processor
	 wi.Add(1)
	 go processAddress(c_addr, c_out)

	//Launch a single output writer
	 wo.Add(1)
	 go outputWriter(c_out)

	//Parse CIDRs and feed IPs to the input channel
	for _, cidr := range flag.Args() {
		AddressesFromCIDR(cidr, c_addr)
	}

	//Close the cidr input channel
	 close(c_addr)

	// Wait for the input feed to complete
	wi.Wait()

	// Wait for pending probes
	waitProbes()


	// Close the output handle
	close(c_out)

	// Wait for the output goroutine
	wo.Wait()
}
