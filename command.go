package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var myPfsCmd = &cobra.Command{
	Use:   "mypfs",
	Short: "My personal file server",
	Long: `A personal file server for sharing files with and receiving files 
from people on your network.
Complete documentation is available at http://github.com/joncrlsn/mypfs`,
	Run: func(cmd *cobra.Command, args []string) {
		action = "up/down"
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Allow file uploads to the current directory",
	Long: `Starts web server that only allows uploading of files to 
the current directory.  No downloading is allowed`,
	Run: func(cmd *cobra.Command, args []string) {
		action = "up"
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Allow file downloads from the current directory",
	Long: `Starts web server that only allows downloading of files from 
the current directory.  No uploading is allowed`,
	Run: func(cmd *cobra.Command, args []string) {
		action = "down"
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  `Displays the version number for mypfs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mypfs (My Personal File Server) version", version)
		action = "version"
	},
}

func init() {
	myPfsCmd.AddCommand(uploadCmd)
	myPfsCmd.AddCommand(downloadCmd)
	myPfsCmd.AddCommand(versionCmd)
	myPfsCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "port number to listen on")
	myPfsCmd.PersistentFlags().Int64VarP(&timeoutMinutes, "timeout", "t", 10, "number of minutes to leave this running")
}
