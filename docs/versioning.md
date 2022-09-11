#### Version Control

*Definition: the task of keeping a software system consisting of many versions and configurations well organized.*

**Best Practices:**
- Always use version control
- All of your source code should be versioned
    - Ex. A new developer should be able to sit down, clone a repo, and run the code with no problem.
- Commit small and often
- Use Descriptive commit messages
- Don't commit broken code.
- Branches should be simple and fit the project needs.
    - the mechanisms to handle changes should be small and simple
- Use commit hooks to enforce quality
- Be carful with secrets (this is when a .gitignore file is very helpful)

**Main Tools**
- Git
- GitHub

**Useful Commands**
```
#To see the commit history for a repository:
$ git log

#To see what changes are available
$ git status

#To see the actual changes in a commit 
$ git diff
```