pipeline {
  agent any
  stages {
    stage('sleep10') {
      parallel {
        stage('sleep') {
          steps {
            sleep 10
          }
        }

        stage('sleep20') {
          steps {
            sleep 20
          }
        }

      }
    }

    stage('end') {
      parallel {
        stage('end') {
          steps {
            sleep 2
          }
        }

        stage('end3') {
          steps {
            sleep 3
          }
        }

      }
    }

    stage('all end') {
      steps {
        sleep 1
      }
    }

  }
}