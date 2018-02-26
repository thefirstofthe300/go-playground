pipeline {
  agent any
  stages {
    stage('Run tests') {
      steps {
        sh '''#! /bin/bash

go test ./...'''
      }
    }
  }
  tools {
    go 'Go 1.8'
  }
}