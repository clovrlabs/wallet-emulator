// +build staging

package helpers

func CopyConfigFiles() {
	Mkdir("data")
	CopyFile("./conf/breez.staging.conf", "./data/breez.conf")
	CopyFile("./conf/lnd.staging.conf", "./data/lnd.conf")
}
