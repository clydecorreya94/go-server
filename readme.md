**To run the code**

    # go run main.go

Use browser - URL

    main page - localhost:8080
    form page - localhost:8080/form.html
    hello page - localhost:8080/hello

**To generate executable file**

    # go build 

**To run executable file**
    # ./go-server

**Issues and debugging**

*Starting server at post 8080*
*listen tcp :8080: bind: address already in use*
*exit status 1*

If the above error return while running the main.go, use the below script

    # sudo kill -9 `sudo lsof -t -i:8080`



