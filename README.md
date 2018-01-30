# fakecc
A fake Cloud Controller server that response the specified "information" endpoints

# Install:
Requires GoLang to be installed:
```
$ go get -u github.com/XenoPhex/fakecc
```

# How to use:
First start the server (this can be backgrounded with an `&` at the end of the command):
```
$ fakecc --disable-ccv3 
API URL: http://localhost:8080
```

With the CF CLI, target the server:
```
$ cf api --skip-ssl-validation http://localhost:8080
Setting api endpoint to http://localhost:8080...
Warning: Insecure http API endpoint detected: secure https API endpoints are recommended
OK

api endpoint:   http://localhost:8080
api version:    2.101.0
Not logged in. Use 'cf login' to log in.

$ cf tasks foo
This command requires CF API version 3.0.0 or higher.
FAILED
```
