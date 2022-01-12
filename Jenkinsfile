pipeline {
  agent any
  stages {
    stage('Source') {
      steps {
        git(url: 'https://github.com/hotttao/goalgo', branch: 'master')
      }
    }
    stage('SonarQube analysis') {
<<<<<<< HEAD
      steps {
        withSonarQubeEnv('SonarQube') {
          sh './gradlew sonarqube'
=======
        steps {
            withSonarQubeEnv('SonarQube') {
                sh "./gradlew sonarqube"
            }
        }
    }
    stage("Quality gate") {
        steps {
            waitForQualityGate abortPipeline: true
>>>>>>> d95d5de (change jenkinsfile)
        }

      }
    }
  }
}