#### Continuous Integration 

**Three Main Categories**

1. Open Source (Jenkins)
    - you host it 
    - configure it
    - maintain it
2. Commercial
    - may be hosted on-premise
    - typically expesive
    - come with alot of support
3. CI as a Service (CircleCI, GitHub Actions)
    - may be hosted in the cloud
    - pay as you go model may help with costs
    - more support than with open source

**Best Practices**
- Start with a clean environment
    - VM
    - container
    - virtual env
- Build to pass the "coffee test"
    - The coffee test is the idea that the amount of time it takes from code commit to receive results should be at most as long as it takes a person to get a cup of coffee (this is probably no more than 5 minutes assuming you make a fresh pot)

**CI Culture**
- Run tests locally before committing
- Don't commit code to broken builds unless it fixes the build
- Don't leave the build broken
- Don't remove failing tests 

**Notifications**
Notifications help mantain transparancy and communication for the development process. For example, sending automatic notifications to a teams channel when a commit fails a test or getting a notification when committed code hasn't been deployed for over an hour. 


**Example**

GitHub actions let you build custom CI solutions using a yaml file stored in your repository. There are many benefits to using GitHub actions including low costs, a large marketplace with pre-built solutions, and minimal development time. 

Since GitHub actions are part of your repository, you can seamlessly add them to part of your development process. 

#### The Yaml file
The yaml file is a markdown file that contains the instructions for your action. 

The basic structure of a workflow is:
- trigger: when should the workflow run
- jobs: what you want the workflow to do
    - runs on: the type of machine you want to jobs to run on
    - steps: the actual actions the workflow will take

The go.yaml file in the .github\workflows directory shows a simple example that builds all binaries and runs any tests in the repository.

#### Building Your Own Actions
While there are many actions already available in the GitHub marketplace, you can build your own custom actions using Javascript/Typescript or using docker and any language you chose. 

References: [Custom GitHub Actions][def]

[def]: https://docs.github.com/en/actions/creating-actions/about-custom-actions