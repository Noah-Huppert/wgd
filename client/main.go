package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Noah-Huppert/wgd/client/rpc"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/Noah-Huppert/goconf"
	"github.com/Noah-Huppert/golog"
	"github.com/vishvananda/netlink"
	"google.golang.org/grpc"
	grpcCredentials "google.golang.org/grpc/credentials"
)

// Wireguard service configuration.
type Config struct {
	// Name of application. Will be presented to users as the identity of
	// your VPN.
	AppName string `default:"Wireguard VPN" validate:"required"`

	// RPC configuration.
	RPC struct {
		// ServerAddr is the address of the RPC server.
		ServerAddr string `default:"127.0.0.1:6000" validate:"required"`

		// CACertificateFile is the certificate of the certficate authority
		// which signed the server's certificates.
		CACertificateFile string `default:"../rpc/dev-ca.cert" validate:"required"`
	}
}

// GUI window. All logic related to laying out the GUI as well
// as any side effect actions occur here.
type Window struct {
	// Context.
	Ctx context.Context

	// Logger.
	Logger golog.Logger

	// Configuration.
	Config Config

	// Data which will be displayed or effect the display of
	// the GUI.
	State GUIState

	// Bus on which longer running non-blocking tasks can signal their results.
	Bus chan WindowEvent

	// Fonts loaded for use in layouts. Pointer bc it will be nil until the
	// fonts are loaded.
	Fonts *WindowFonts

	// Client for registry service.
	RegistryClient rpc.RegistryClient
}

// Fonts for use in Window layouts.
type WindowFonts struct {
	// Zilla Slab Regular
	ZillaSlabRegular imgui.Font
}

// Window task bus events.
type WindowEvent interface {
	// Commit executes the side effect of the event on the window. This could
	// be a change to a label's text or a tweak to the state. Events should
	// perform their changes to the window here. This method should not block
	// and exit as quickly as possible.
	Commit(w *Window) error
}

// Window event which displays an error to the user.
type ErrorEvent struct {
	// User friendly error.
	error string
}

// Create an error event.
func NewErrorEvent(e string) ErrorEvent {
	return ErrorEvent{
		error: e,
	}
}

// Add the error to the window's state.
func (e ErrorEvent) Commit(w *Window) error {
	w.State.Errors = append(w.State.Errors, e.error)
	return nil
}

// State of GUI components. Data in here will effect how or what is shown
// in the GUI.
type GUIState struct {
	// Errors which have occurred and are being shown to the user.
	Errors []string

	// Indicates the how the process of loading the machine's interfaces
	// is going.
	LoadWgIfacesStatus WgIfaceStatus

	// List of Wireguard interfaces and their state
	WgIfaces []WgIfaceState
}

// Wireguard interface states.
type WgIfaceState struct {
	// Name of interface.
	Name string

	// Status of interface.
	Status WgIfaceStatus
}

// Indicates the status of a Wireguard interface or of a process related to
// an interface.
type WgIfaceStatus string

const (
	// System is currently loading the interface's state.
	WgIfaceLoading WgIfaceStatus = "Loading"

	// Interface is being setup.
	WgIfaceSettingUp WgIfaceStatus = "Setting Up"

	// Ready.
	WgIfaceReady WgIfaceStatus = "Ready"

	// An error occurred.
	WgIfaceError WgIfaceStatus = "Error"
)

// New GUI state.
func NewGUIState() GUIState {
	return GUIState{
		Errors:             []string{},
		LoadWgIfacesStatus: WgIfaceLoading,
		WgIfaces:           []WgIfaceState{},
	}
}

// Initializes a new Window
func NewWindow(ctx context.Context, baseLogger golog.Logger) (*Window, error) {
	logger := baseLogger.GetChild("window")

	// Load configuration
	cfgLdr := goconf.NewLoader()
	cfgLdr.AddConfigPath("/etc/wgd/*")
	config := Config{}
	if err := cfgLdr.Load(&config); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %s", err)
	}

	// Create registry service client
	caCertBytes, err := ioutil.ReadFile(config.RPC.CACertificateFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read server certificate authority "+
			"certificate file: %s", err)
	}

	tlsCertPool := x509.NewCertPool()
	if !tlsCertPool.AppendCertsFromPEM(caCertBytes) {
		return nil, fmt.Errorf("failed to add server certificate authority "+
			"to a x509 cert pool: %s", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            tlsCertPool,
	}

	grpcConn, err := grpc.Dial(config.RPC.ServerAddr,
		grpc.WithTransportCredentials(grpcCredentials.NewTLS(tlsConfig)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %s", err)
	}

	registryClient := rpc.NewRegistryClient(grpcConn)

	return &Window{
		Ctx:            ctx,
		Logger:         logger,
		Config:         config,
		State:          NewGUIState(),
		Bus:            make(chan WindowEvent),
		Fonts:          nil,
		RegistryClient: registryClient,
	}, nil
}

// Load fonts for use. Populates .Fonts field.
func (w *Window) LoadFonts() {
	cwd, err := os.Getwd()
	if err != nil {
		w.Logger.Fatalf("Failed to get current working directory for use "+
			"in loading fonts: %s", err)
	}

	fonts := g.Context.IO().Fonts()

	zillaSlabFont := fonts.AddFontFromFileTTFV(
		filepath.Join(cwd,
			"./fonts/zilla-slab/ttf/ZillaSlab-Regular.ttf"),
		24,
		imgui.DefaultFontConfig,
		fonts.GlyphRangesDefault(),
	)

	w.Fonts = &WindowFonts{
		ZillaSlabRegular: zillaSlabFont,
	}
}

// Display the window.
func (w *Window) Display() {
	// Setup window
	masterWindow := g.NewMasterWindow(w.Config.AppName,
		400, 200, 0, w.LoadFonts)

	w.Logger.Debug("Running initialization logic")
	w.Run(w.LoadInterfaces)

	w.Logger.Debug("Running main event loop")
	masterWindow.Main(w.EventLoop)
}

// Run a background task that can pass its result to the UI via WindowEvents.
func (w *Window) Run(task func() WindowEvent) {
	go func() {
		resEvent := task()
		w.Bus <- resEvent
	}()
}

// Load the machine's interfaces and reflect it in the GUI.
func (w *Window) LoadInterfaces() WindowEvent {
	w.Logger.Debug("Loading interfaces")

	links, err := netlink.LinkList()
	if err != nil {
		w.Logger.Errorf("failed to list links: %s", err)
		return NewErrorEvent("Failed to the VPN's status.")
	}

	ifaces := []WgIfaceState{}
	for _, link := range links {
		a := link.Attrs()
		ifaces = append(ifaces, WgIfaceState{
			Name:   a.Name,
			Status: WgIfaceReady,
		})
		w.Logger.Debugf("type=%s name=%s", link.Type(), a.Name)
	}

	return IfacesLoadedEvent{
		WgIfaces: ifaces,
	}
}

// Occurs when interfaces are loaded.
type IfacesLoadedEvent struct {
	// Loaded interfaces.
	WgIfaces []WgIfaceState
}

// Save new interaces and set the loading state to ready.
func (e IfacesLoadedEvent) Commit(w *Window) error {
	w.State.WgIfaces = e.WgIfaces
	w.State.LoadWgIfacesStatus = WgIfaceReady
	return nil
}

// Main event loop for window
func (w *Window) EventLoop() {
	// Receive any window events
	select {
	case event := <-w.Bus:
		err := event.Commit(w)
		if err != nil {
			w.Logger.Errorf("Failed to run \"%#v\" event commit: %s",
				event, err)
		}
	default:
		break
	}

	// Set default font
	g.PushFont(w.Fonts.ZillaSlabRegular)

	// Setup top menu bar
	layout := g.Layout{
		g.MenuBar(g.Layout{
			g.Menu("Debug", g.Layout{
				g.MenuItem("List interfaces", nil),
			}),
		}),
	}

	// Display any errors
	for _, err := range w.State.Errors {
		layout = append(layout, g.Line(
			g.Label(err),
		))
	}

	// Display status of interfaces
	layout = append(layout, g.Line(
		g.Label(string(w.State.LoadWgIfacesStatus)),
	))

	if w.State.LoadWgIfacesStatus == WgIfaceReady {
		for _, wgIface := range w.State.WgIfaces {
			layout = append(layout, g.Line(
				g.Label(fmt.Sprintf("(%s) %s", wgIface.Status,
					wgIface.Name)),
			))
		}
	}

	g.PopFont()

	// Make window with layout
	g.SingleWindowWithMenuBar(w.Config.AppName, layout)
}

func main() {
	ctx := context.Background()

	baseLogger := golog.NewLogger("wgd")
	baseLogger.Debug("Starting")

	window, err := NewWindow(ctx, baseLogger)
	if err != nil {
		baseLogger.Fatalf("Failed to initialize window: %s", err)
	}
	baseLogger.Debug("Created window")
	window.Display()
}
