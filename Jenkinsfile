pipeline {
  agent any
  stages {
    stage('Checkout') {
      steps {
        git(url: 'https://github.com/viswanathch999/sample-java-web.git', branch: 'master', credentialsId: 'viswanathch999')
      }
    }
    stage('Initialize') {
      steps {
        bat '''
                    export PATH=C:\Program Files\apache-maven-3.5.4
                    export M2_HOME=C:\Program Files\apache-maven-3.5.4
                '''
      }
    }
    stage('Build') {
      steps {
        echo 'build'
        sh 'mvn clean test'
      }
    }
    stage('Test') {
      parallel {
        stage('Test') {
          steps {
            echo 'Tests'
          }
        }
        stage('Junit') {
          steps {
            echo 'junit tests'
            sh 'mvn test -B -Dmaven.javadoc.skip=true -Dcheckstyle.skip=true'
          }
        }
        stage('cucumber') {
          steps {
            echo 'cucumber test cases'
          }
        }
      }
    }
    stage('Dev') {
      steps {
        echo 'Dev env'
      }
    }
    stage('Staging') {
      steps {
        echo 'stage env'
      }
    }
    stage('Production') {
      steps {
        echo 'prod env'
      }
    }
  }
}
