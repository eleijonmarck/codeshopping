#!/usr/bin/env groovy

stageUrl = 'localhost:8080'
stage 'CI'
node {
     git branch: 'master',
     url: 'https://github.com/eleijonmarck/codeshopping.git'
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
}
