# Codeshopping
A webshop purely focused on programming / code / machine learning / math / computer science

## projects that have inspired the structure of this project

* [generic shopping cart](https://github.com/fernandez14/go-cart)
* [goddd](https://github.com/marcusolsson/goddd/blob/master/main.go)

Continuous Deployment
1. We SCM-poll the github repo or POST changes via a webhook
2. Project is built via ```make build```
3. Tests are run ```make test```
 * all tests log whatever has happened that could be wrong
4. We fetch the clientresources JS/CSS/Images/Etc
5. All files are deployed if test are green
6. The host restarts and the new version of the site is up and running
