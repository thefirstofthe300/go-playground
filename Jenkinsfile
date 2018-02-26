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
}