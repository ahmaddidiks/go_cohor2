package main

import "gin/routers"

func main() {
    var PORT = ":9999"

    routers.StartServer().Run(PORT)
}