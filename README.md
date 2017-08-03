<p align="center">
    <p align="center"><strong><font size="6">Gostagram</font></strong></p>
    <p align="center">Unofficial and easy to use instagram client for go.</p>
    <p align="center">
      <a href="https://github.com/leoxnidas/gostagram/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/leoxnidas/gostagram/all.svg?style=flat-square"></a>
      <a href="https://github.com/leoxnidas/gostagram/"><img alt="issues" src="https://img.shields.io/github/downloads/leoxnidas/gostagram/total.svg?style=flat-square"/></a>
      <a href="/LICENSE.text"><img alt="license" src="https://img.shields.io/github/license/leoxnidas/gostagram.svg?style=flat-square"/></a>
      <a href="https://github.com/leoxnidas/gostagram"><img alt="license" src="https://img.shields.io/github/contributors/leoxnidas/gostagram.svg?style=flat-square"/></a>
      <a href="https://github.com/leoxnidas/gostagram"><img alt="Powered By: gostagram" src="https://img.shields.io/badge/powered%20by-gostagram-green.svg?style=flat-square"></a>
     </p>
</p>

---

**Warning!**
gostagram is in alpha release, is not stable. Use it on your own risk.

### Quick Start.

**First step**

Go to instagram developer [website](https://www.instagram.com/developer/)
and create a developer account, then register a new instagram client.

**Download and Installation**
```text
go get github.com/leoxnidas/gostagram
```

**Usage**

Basic example, using an client to
consume an instagram endpoint.

```go
package main

import (
    "fmt"
    "github.com/leoxnidas/gostagram"
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
    "github.com/leoxnidas/gostagram"
)

func main() {
    client := gostagram.NewClient("access_token")
    client.SetSignedRequest(true)
    client.SetClientSecret("client secret")
    
    
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
