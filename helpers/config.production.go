// +build !develop,!staging

package helpers

func CopyConfigFiles() {
	Mkdir("data")
	CopyFile("./conf/breez.production.conf", "./data/breez.conf")
	CopyFile("./conf/lnd.production.conf", "./data/lnd.conf")
}
