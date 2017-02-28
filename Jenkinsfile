#!/usr/bin/env groovy

node {
    checkout scm
        stage('Build') {
                echo 'Building......'
                echo 'wtf.....'
        }
        stage('Test') {
            make test
        }

        stage('Deploy') {
                echo 'Deploying....'
        }
}
