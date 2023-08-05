# FROM quay.io/travisci/travis-python
FROM travisci/ci-python:packer-1490914243
MAINTAINER "Ashwin Vishnu Mohanan <avmo@kth.se>"
LABEL Description="A docker instance to debug Travis Python builds."

USER travis
ARG HOME=/home/travis
WORKDIR /home/travis

RUN bash -lc "rvm install 2.3.0"
RUN bash -lc "rvm use 2.3.0"
RUN bash -lc "gem install travis --no-doc --no-ri"
RUN bash -lc "travis version"
RUN bash -lc "gem install bundler"

RUN git clone https://github.com/travis-ci/travis-build.git ~/.travis/travis-build
WORKDIR ~/.travis/travis-build
RUN bash -lc "bundle install --gemfile ~/.travis/travis-build/Gemfile"
RUN bash -lc "bundler add travis"
RUN bash -lc "bundler binstubs travis"
RUN mkdir build
RUN echo "export TRAVIS_BUILD_APT_PACKAGE_WHITELIST=https://raw.githubusercontent.com/travis-ci/apt-package-whitelist/master/ubuntu-trusty" >> .bashrc 
RUN echo "export TRAVIS_BUILD_APT_SOURCE_WHITELIST=https://raw.githubusercontent.com/travis-ci/apt-source-whitelist/master/ubuntu.json" >> .bashrc
RUN echo "\
if [ -f ~/.bashrc ] ; then \
   source ~/.bashrc \
fi" >> .bash_profile

WORKDIR /home/travis/build
CMD /bin/bash
