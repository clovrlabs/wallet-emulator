// +build develop

package helpers

func CopyConfigFiles() {
	Mkdir("data")
	CopyFile("./conf/breez.develop.conf", "./data/breez.conf")
	CopyFile("./conf/lnd.develop.conf", "./data/lnd.conf")
}
