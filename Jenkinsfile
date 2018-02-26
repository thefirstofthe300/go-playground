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

printenv
go version'''
      }
    }
  }
}
