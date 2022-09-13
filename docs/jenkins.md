#### Jenkins tutorial

*Definition: Jenkins is an open source automation server*

Jenkins facilitates continuous integration and continuous delivery in software projects by automating parts related to build, test, and deployment. This makes it easy for developers to continuously work on the betterment of the product by integrating changes to the project.

Jenkins automates the software builds in a continuous manner and lets the developers know about the errors at an early stage.


**Key Terms**
1. Project/Job -  a collection of tests you want jenkins to manage
2. Build - one run of a project (build a project/start a build)
3. Build Step - a task inside a project
4. Build Trigger - criteria for starting a build
5. Plugin - an extension to the base jenkins functionality

**Installing Jenkins**

Using Docker

Docker is a great way to start from a clean slate and keep our machine seperate from our development environment. You can 
always install locally but using docker is really simple.

For this I'll be using the jenkins 2.36.1.1lts-alpine image which runs using Alpine Linux.

Commands to get started:

```bash
#get the docker image
docker pull jenkins/jenkins:2.361.1-lts-alpine

#run the docker image
docker run -d --publish 8080:8080 --volume jenkins_home:/var/jenkins_home --name jenkins jenkins/jenkins:2.361.1-lts-alpine 

#for initial login only
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPasswordbas
```
And that's it! You can now access Jenkins on localhost:8080