# youarehere-eyes

A cross-platform desktop app for setting your desktop background based on your geolocation. Like [Satellite Eyes](https://github.com/tomtaylor/satellite-eyes) but Mac OS, Windows, and Linux.

## What is this?

For years, I've been a huge fan of Tom Taylor's [Satellite Eyes](https://github.com/tomtaylor/satellite-eyes) a beautiful MacOS app that sets your desktop background to raster map tiles based on your location. On a number of occasions, I've thought of switching over to other OSes, only to come back mostly for Satellite Eyes. So this is my attempt to update it for a modern-ish era.

## Does it work

No. Not in the least. It's basically just this file.

## Roadmap

### Proof of concept

Cross platform CLI tool that:

1. geolocates based on wifi locations (and otherwise falls back to geoip)
2. Fetches a static set of map tiles with that location in the center
3. Stitches them together into a single large image that matches screen resolution
4. Sets that image as the desktop background

Target Priority:

1. Mac OS
2. Ubuntu
3. Windows 10

### Muliple Map Mode

* Enable adding multiple tile layers and switching them out.
* Bonus [Use go-rasterzen](https://github.com/whosonfirst/go-rasterzen) for the Mapzen styles
* Set different zoom levels

### Multiple Monitor Mode

1. Detect number of monitors connected
2. Determine which monitor is "default" and relative placement of other displays and set the center of the default monitor  as the centerpoint.
3. Fetch and tiles to cover all monitors (at each monitor's default resolution)

### GUI
Use a library like [systray](https://github.com/getlantern/systray) to create a cross-platform UI for running cross platform in system tray.

* [ ] When running in system tray should fetch new locations every time computer wakes from sleep or every n minutes (configurable by user)
* [ ] Should allow selecting map layer and zoom layer
* [ ] Should allow custom layers


## Useful Utilities I might use:

* https://github.com/getlantern/systray
* https://github.com/whosonfirst/go-rasterzen
* [Here.xyz Golang Wifi Tutorial](https://developer.here.com/blog/tracking-a-raspberry-pi-with-wlan-and-golang-then-displaying-the-results-with-here-xyz)
* https://github.com/schollz/wifiscan

