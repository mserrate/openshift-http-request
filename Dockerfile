FROM scratch
MAINTAINER Marçal Serrate <mserrate@gmail.com>
COPY bin/openshift-http-request /openshift-http-request
EXPOSE 8080 8888
ENTRYPOINT ["/openshift-http-request"]