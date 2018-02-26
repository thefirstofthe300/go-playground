pipeline {
  agent any
  
  tools {
    go 'Go 1.8'
  }
  stages {
    stage ('Initialize') {
            steps {
                sh '''
                    echo "PATH = ${PATH}"
                    echo "GOROOT = ${GOROOT}"
                    printenv
                ''' 
            }
        }
    stage('Run tests') {
      steps {
        sh '''#! /bin/bash

go test ./...'''
      }
    }
  }
}
