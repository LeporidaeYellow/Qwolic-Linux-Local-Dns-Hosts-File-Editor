package main

import (
	"flag"
	"hostsEditor/config"
	service "hostsEditor/web/controller"
)

func main() {
	tempfile := flag.String("tf", "tempFile", "The name of temporary file")
	backupfiles := flag.String("bf", "files/", "The folder of backup files")
	filehosts := flag.String("fh", "/etc/hosts", "The file of hosts")
	offsetfilehosts := flag.Int("ofh", 4, "The number of read offset for hosts file")
	flag.Parse()

	config.VarTempFile = *tempfile
	config.VarBackupFiles = *backupfiles
	config.VarFileHost = *filehosts
	config.VarOffsetFileHost = *offsetfilehosts
	service.HostService()
}
