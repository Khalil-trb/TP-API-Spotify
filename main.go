package main

import (
    "log"
    "net/http"
    "spotify/router"
)

func main() {

    mux := http.NewServeMux()

   
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/templates/index.html", http.StatusFound)
    })

 
    mux.Handle("/templates/",
        http.StripPrefix("/templates/",
            http.FileServer(http.Dir("./templates")),
        ),
    )

   
    mux.Handle("/style/",
        http.StripPrefix("/style/",
            http.FileServer(http.Dir("./style")),
        ),
    )

    
    mux.Handle("/img/",
        http.StripPrefix("/img/",
            http.FileServer(http.Dir("./img")),
        ),
    )

    
    mux = router.NewWithMux(mux)

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
