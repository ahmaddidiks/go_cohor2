package main

import "go_cohort_2/sesi5/gin_framwork/routers"

func main() {
    var PORT = ":9999"

    routers.StartServer().Run(PORT)
}