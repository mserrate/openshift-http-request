Openshift HTTP request
----------------------

Sample based in: [Openshift samples](https://raw.githubusercontent.com/openshift/origin/master/examples/hello-openshift)

This example will serve an HTTP response and will show the HTTP request in the log: useful to test and check HTTP security headers, etc.

    $ oc create -f openshift-http-request.yml


To test from external network, you need to create router. Please refer to [Running the router](https://github.com/openshift/origin/blob/master/docs/routing.md)

If you need to rebuild the image:

    $ env GOOS=linux GOARCH=386 go build # cross-platform build
    $ mv openshift-http-request bin
    $ docker build -t docker.io/mserrate/openshift-http-request .

