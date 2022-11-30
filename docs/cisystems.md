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
3. CI as a Service (CircleCI)
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