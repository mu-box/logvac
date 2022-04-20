# Contributor: Hennik Hunsaker <hennikhunsaker@microbox.cloud>
# Maintainer: Hennik Hunsaker <hennikhunsaker@microbox.cloud>
pkgname=logvac
pkgver=0.1.1
pkgrel=0
pkgdesc="Simple, lightweight, api-driven log aggregation service with realtime push capabilities and historical persistence."
url="https://github.com/mu-box/logvac"
arch="all"
license="MIT"
depends=""
makedepends="go git bash"
checkdepends=""
install=""
subpackages=""
source=""
srcdir="/tmp/abuild/logvac"
builddir=""

build() {
	go get -t -v .
	go install github.com/mitchellh/gox@latest
	export PATH="$(go env | grep GOPATH | sed -E 's/GOPATH="(.*)"/\1/')/bin:${PATH}"
	./scripts/build.sh
}

check() {
	# Replace with proper check command(s)
	:
}

package() {
	install -m 0755 -D ./build/logvac "$pkgdir"/sbin/logvac
}
