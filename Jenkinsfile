#!/usr/bin/env groovy

node {
    checkout scm
        stage('Build') {
                echo 'Building......'
                echo 'wtf.....'
        }
        stage('Test') {
            echo "Running make test"
            sh "make test"
        }

        stage('Deploy') {
                echo 'Deploying....'
        }
}
