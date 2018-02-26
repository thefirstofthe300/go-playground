pipeline {
  agent any
  tools {
    go 'Go 1.8'
  }
  stages {
    stage('Run tests') {
      steps {
        sh '''#! /bin/bash

go version'''
      }
    }
  }
}
