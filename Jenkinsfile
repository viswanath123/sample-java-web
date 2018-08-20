pipeline {
  agent any
  stages {
    stage('Checkout') {
      steps {
        git(url: 'https://github.com/santosh52krishna/os-sample-java-web.git', branch: 'master', credentialsId: 'santosh52krishna')
      }
    }
    
     stage ('Initialize') {
            steps {
                sh '''
                    export PATH=/home/ubuntu/apache-maven-3.5.3/bin
                    export M2_HOME=/home/ubuntu/apache-maven-3.5.3
                '''
            }
        }
    
    stage('Build') {
      steps {
        echo 'build'
        sh '''mvn clean -Dmaven.test.failure.ignore=true install

echo "test build"'''
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
