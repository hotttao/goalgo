pipeline {
  agent any
  stages {
    stage('Source') {
      steps {
        git(url: 'https://github.com/hotttao/goalgo', branch: 'master')
      }
    }

    stage('SonarQube analysis') {
            steps {
                withSonarQubeEnv('SonarQube') {
                    sh "./gradlew sonarqube"
                }
            }
        }
    
    stage("Quality gate") {
        steps {
            waitForQualityGate abortPipeline: true
        }
    }

  }
}