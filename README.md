# mypfs
A "personal fileserver" for sharing files with (and receiving files from) other computers on your network. More specifically, it is a small web server that exposes the files in the current directory and/or allows uploads to the same directory.

### why?
* Me: Sarah, can you send me that 100MB zip of those log files?
* Sarah: Sure, I'll attach it to an email
* Me: Um, our email system doesn't allow attachments over 10MB
* Sarah: OK, do you have a flash drive?
* Me: No, that's such a pain.
* (I start mypfs)
* Me: Here, open your browser to `http://<my-internal-ip-address>:8080` and upload the file
* (30 seconds later)
* Sarah: Wow, that was sooo easy!

### features
1. mypfs sports an HTTP web interface to upload and/or download files to the current directory
1. runs on the command-line
1. during startup you can specify upload, download, or accept the default of both
1. server will run for 10 minutes by default, then exit -- (this is a security feature which gives enough time to exchange files and protect you if you forget to shut it off)

### examples
* `mypfs version`
* `mypfs --timeout=5 --port=8888` (both upload and download)
* `mypfs -t5 -p8888 download`
* `mypfs upload`

### how to install and run
1. download executable for your platform ( [windows](https://github.com/joncrlsn/mypfs/raw/master/bin-win/mypfs.exe "Windows"), [osx](https://github.com/joncrlsn/mypfs/raw/master/bin-osx/mypfs "OSX"), [linux](https://github.com/joncrlsn/mypfs/raw/master/bin-linux/mypfs "Linux") )
1. place executable somewhere in your path
1. navigate to the directory with files you want to share 
1. run `mypfs`
1. share URL with person you are sharing with. i.e.  http://\<my-ip-address\>:8080

### safe use
1. mypfs will work over the internet only if your computer has a public IP address or you have port-forwarding setup on your router, however direct internet use is discouraged
1. run mypfs in a small directory, never the root or home directory
1. avoid use on a public network (like a coffeeshop or the internet)
1. avoid extending the timeout unless you totally trust your network
1. shut it down as soon as you've exchanged files

### cool things to improve on
1. modify FileServer to log when someone downloads a file
1. on the web pages, show time remaining until server shuts down
1. instead of exiting after timeout, show a "timeout" page
1. limit uploads to a configurable amount 
1. after startup, display HTTP URL to copy and share
1. add parameter with directory to be served  i.e. `mypfs upload /tmp/share`
