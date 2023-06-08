#  SALVA


[![](https://img.shields.io/github/actions/workflow/status/yahyrparedes/salva/test.yml?branch=main&longCache=true&label=Test&logo=github%20actions&logoColor=fff)](https://github.com/yahyrparedes/salva/actions?query=workflow%3ATest)
[![Go Reference](https://pkg.go.dev/badge/github.com/yahyrparedes/salva.svg)](https://pkg.go.dev/github.com/yahyrparedes/salva)
[![Go Report Card](https://goreportcard.com/badge/github.com/yahyrparedes/salva)](https://goreportcard.com/report/github.com/yahyrparedes/salva) 


## Overview
- A quick easy start to increase your development speed.
- Try 'salva' and generate files for your API and increment your development sped
- Create Controller, Model, Router, Service, Repository and Testing very easy.

### Recommendation
It is recommended to use this template **[Salva Template](https://github.com/yahyrparedes/salva-template)** or guide folder structure for the best saving experience.

### Install
```bash 
go get -u github.com/spf13/cobra/cobra
```

### How to use

#### Generate Controller for User
```bash 
salva controller user
```

#### Generate Route for User
```bash 
salva route user
```

#### Generate Service for User
```bash  
salva service user
```

#### Generate Model for User
```bash 
salva model user
```

#### Generate Repository for User
```bash 
salva repository user
```

#### Generate All for User
```bash 
salva magic user
```

#### Generate specific for User
```bash 
salva magic user --router --controller --model
```
or 
```bash 
salva magic user -r -c -m
```



### Go Archetype (template)

#### [Salva Template](https://github.com/yahyrparedes/salva-template) 
 https://github.com/yahyrparedes/salva-template 


> ### SALVA CLI is created to streamline development work thinking of a robust, simple and understandable architecture.