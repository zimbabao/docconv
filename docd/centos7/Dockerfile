FROM centos:7

MAINTAINER Hamish Ogilvy

RUN yum -y install \
 zip \
 poppler-utils \
 wv \
 unrtf \
 tidy \
 lynx \
 && yum clean all
 
EXPOSE 8888

COPY docd /
CMD ["/docd"]
