pipeline {
  agent any
  tools {
    go 'Go 1.8'
  }
  environment {
    PATH = "$GOROOT/bin:$PATH"
  }
  stages {
    stage('Run tests') {
      steps {
        sh '''#! /bin/bash

/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.8/bin/go version'''
      }
    }
  }
}
