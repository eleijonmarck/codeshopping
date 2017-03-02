#!/usr/bin/env groovy

node {
    checkout scm
        stage('Build') {
                echo 'Building......'
                echo 'wtf.....'
        }
        stage('Test') {
            echo "Running make test"
            sh "go test -race -v $(go list ./... | grep -v /vendor/)"
            echo "ran the tests"
        }

        stage('Deploy') {
                echo 'Deploying....'
        }
}
