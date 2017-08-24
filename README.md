<p align="center">
    <h3 align="center"><strong>Gostagram</strong></h3>
    <p align="center">Unofficial and easy to use instagram client for go.</p>
    <p align="center">
      <a href="http://godoc.org/github.com/ljesparis/gostagram"><img alt="gostagram documentation" src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"/></a>
      <a href="https://github.com/leoxnidas/gostagram/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/leoxnidas/gostagram/all.svg?style=flat-square"></a>
      <a href="/LICENSE.text"><img alt="license" src="https://img.shields.io/github/license/leoxnidas/gostagram.svg?style=flat-square"/></a>
      <a href="https://github.com/leoxnidas/gostagram"><img alt="Powered By: gostagram" src="https://img.shields.io/badge/powered%20by-gostagram-green.svg?style=flat-square"></a>
     </p>
</p>

---

### Quick Start.

**Create Instagram Client**

Go to instagram developer [website](https://www.instagram.com/developer/)
and create a developer account, then register a new instagram client.

**Implement Oauth2 Protocol**

Get the access token, implementing oauth2 [authorization protocol](https://en.wikipedia.org/wiki/OAuth).
I do recommmend [oauth2](https://github.com/golang/oauth2) for this job.
Here you can find an [oauth2 example](https://github.com/dorajistyle/goyangi/tree/master/util/oauth2).

**Download and Installation**

```text
go get github.com/ljesparis/gostagram
```

**Usage**

Basic example, using an client to
consume an instagram endpoint.

```go
package main

import (
    "fmt"
    "github.com/ljesparis/gostagram"
)

func main() {
    client := gostagram.NewClient("access_token")
    user, err := client.GetCurrentUser()
    
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(user.Id)
        fmt.Println(user.Username)
        fmt.Println(user.FullName)
    }
}
```

If you want to enable instagram signed requests mode
you have to tell gostagram client, that you want to sig
a request.

```go
package main

import (
    "fmt"
    "github.com/ljesparis/gostagram"
)

func main() {
    client := gostagram.NewClient("access_token")
    client.SetSignedRequest(true)
    client.SetClientSecret("client secret")
    
    
    tags, err := client.SearchTags("golang")
    
    if err != nil {
        fmt.Println(err)
    } else {
        for _, tag := range tags {
          fmt.Println("Tag name: ", tag.Name)
        }
    }
}
```

### Tests.
Before executing **gostagram** tests, please get access token, client secret
and complete every empty variable in all test file, why? well, that way you could
test every method with your own parameters, otherwise multiples errors will be
thrown.

Note: test every method one by one, use the makefile to optimize that
process.

### Support us.
 * [donate](https://www.paypal.me/leoxnidas).
 * [Contribute](https://github.com/leoxnidas/gostagram#contribute).
 * Talk about the project.

### Contribute.
Please use [Github issue tracker](https://github.com/leoxnidas/gostagram/issues)
for everything.
  * Report issues.
  * Improve/Fix documentation.
  * Suggest new features or enhancements.

### License.
gostagram license [MIT](./LICENSE.txt)
