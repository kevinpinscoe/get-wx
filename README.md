# get-wx A very simple open meteo weather parser written in golang

wx is a aeronautical term for weather or weather report.

## Prerequisites

* A Linux or Apple Mac PC. Windows would probably work in cygwin or WSL I didn't try.
* Install the golang compiler and make tools using your package manager or brew for Mac.
* Internet connection

## Install

### Package managers

#### Homebrew (macOS/Linux)

```bash
brew tap kevinpinscoe/homebrew-tap
brew install get-wx
```

#### APT (Debian/Ubuntu)

```bash
curl -sL https://kevinpinscoe.github.io/apt/gpg.key \
  | sudo gpg --dearmor -o /etc/apt/keyrings/kevinpinscoe.gpg

echo "deb [signed-by=/etc/apt/keyrings/kevinpinscoe.gpg] \
  https://kevinpinscoe.github.io/apt stable main" \
  | sudo tee /etc/apt/sources.list.d/kevinpinscoe.list

sudo apt update
sudo apt install get-wx
```

#### DNF (Fedora/RHEL)

```bash
sudo curl -fsSL https://kevinpinscoe.github.io/rpm/kevinpinscoe.repo \
  -o /etc/yum.repos.d/kevinpinscoe.repo
sudo dnf install get-wx
```

### Compile from source

* Clone this repo (git clone)
* run make
* run make install (installs to your home directory under the bin subdirectory). You can edit Makefile to change the location of the install. Become sudo if using a system location.

## Run

* First determine your latitude and longitude in decimal degrees. If you don't know yours visit Google Maps and right click on your location. For privacy You will only require two digits nn.nn or -nn.nn. 
* On first run you will be prompted for your latitude and longitude. Enter them as decimal degrees and not degrees, minuted, seconds or any other format or the retrieve will not work. 
* A config file will be saved in your home directory under the .config subdirectory named get-wx.
