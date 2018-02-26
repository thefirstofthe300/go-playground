pipeline {
  agent any
  tools {
    go 'Go 1.8'
  }
  stages {
    stage('Run tests') {
      steps {
        sh '''#! /bin/bash

set +x
ls /var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.8/bin
go version'''
      }
    }
  }
}
