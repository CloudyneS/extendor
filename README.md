![](img/extendor.svg)

# Cloudyne Extendor: Go Packages!
Extendor is a package manager, currently only with support for Composer but with more to come.
It's designed to be simple, lightweight and to not require a cacaphony of dependencies, databases, caches or other dependencies.

## API Overview
### URLs
Given the base URL of 'https://demo.extendor.com' the following URLs are valid for the repository with the name 'composer'. The type of repository is not included in the URL.
- https://demo.extendor.com/<request-path> (If multi-repo is disabled, you can access all repos directly from the base URL)
- https://demo.extendor.com/composer/<request-path> (Works regardless of multi-repo flag)
- https://composer.demo.extendor.com/<request-path> (Works regardless of multi-repo flag)

### Authentication
There are three types of authentication possible to access or upload packages, Basic Auth, Token or JWT. Authentication can be set on the server or repository level.

#### None
No authentication is required to access packages, uploading is disabled

#### Static Token
A static token set in the configuration file on the server or repository level.
Sent via the Authorization header:

```json
{
    "Authorization": "Bearer <token>"
}
```

#### BASIC Auth
Uses your username and password or personal access token to authenticate.

```json
{
    "Authorization": "Basic <base64-encoded-username:password>"
}

```

#### Dynamic Token (TBD)
Access-controlled token with specific permissions

#### JWT (TBD)
Authenticate with JWT

### General
#### GET/POST /ping
Returns a simple pong message for healthchecks
```json
{
    "message": "pong"
}
```

#### POST /upload/<package-repo>/from-git
Uploads a package containing files from a GIT repository

#### POST /upload/<package-repo>/from-file
Upload an archive containing the package directly

#### POST /upload/<package-repo>/from-url
Fetches the package from a remote URL

#### POST /upload/<package-repo>/mirror
Mirror a remote repository, extendor checks for updates on each pull and serves the newer package if one is present

### Composer
#### /packages.json
Returns a list of endpoints for the client to use
```json
{
    "search" : "https://demo.extendor.com/p2/search.json?q=%query%&type=%type%",
    "metadata-url" : "https://demo.extendor.com/p2/packages/%package%.json",
    "list" : "https://demo.extendor.com/p2/list.json"
}
```

#### /p2/search.json
Searches for packages matching the query and type
```json
{
    "results" : [
        {
            "name": "extendor/package-a",
            "version": "1.0.0",
            "description": "Extendor Example 1",
            "type": "library",
            "license": "MIT",
            "url": "https://demo.extendor.com/p2/packages/extendor/package-a.json",
            "repository": "https://github.com/cloudynes/extendor.git",
            "downloads": 999,
            "favers": 1313
        }
    ]
}
```

#### /p2/list.json
Returns a list of all packages
```json
{
    "packages": {
      "extendor/package-a": {
        "1.0.0": {
            "name": "extendor/package-a",
            "version": "1.0.0",
            "description": "Extendor Example 1",
            "type": "library",
            "license": "MIT",
            "authors": [
                {
                "name": "John Doe",
                "email": "johndoe@example.com"
                }
            ],
            "require": {
                "php": "^8.0"
            }
        }
      }
    }
}
```

#### /p2/packages/%package%.json
Returns the package information including download URL
```json
{
    "name": "extendor/package-a",
    "version": "1.0.0",
    "description": "Extendor Example 1",
    "type": "library",
    "license": "MIT",
    "authors": [
        {
        "name": "John Doe",
        "email": "johndoe@example.com"
        }
    ],
    "require": {
        "php": "^8.0"
    },
    "dist": {
        "url": "https://demo.extendor.com/storage/composer/extendor/package-a-1.0.0.zip",
        "type": "zip"
    }
}