- title : serverless
- description : Introduction into serverless 
- author : Sotirios Mantziaris
- theme : night
- transition : default

***

## serverless

<br />
<br />

### Introduction into serverless with Typescript mainly on AWS

<br />
<br />
Sotirios Mantziaris
<br />
Refreshment Engineer (καφετζής)
<br />
smantziaris@gmail.com
<br />
[@smantziaris](http://www.twitter.com/smantziaris)
<br />
[blog](http://blog.mantziaris.eu)
<br />
[github](https://github.com/mantzas)
<br />
[linkedin](https://gr.linkedin.com/in/mantzas)

***

### Host Evolution?

#### On Premises

> Bare Metal
<br />
> VM
<br />
> Docker

#### Cloud 

> VM
<br />
> Docker
<br />
> serverless

---

### What about speed of delivery (with provisioning)?

#### On Premises

> Bare Metal ~ ***weeks?***
<br />
> VM ~ ***hours-days?***
<br />
> Docker ~ ***minutes-days?***

#### Cloud 

> VM ~ ***hours-days?***
<br />
> Docker ~ ***minutes-days?***
<br />
> serverless ~ ***seconds?***

***

### What is serverless?

***

### Who offers serverless

- [AWS Lambda](https://aws.amazon.com/lambda/)
- [Azure Functions](https://docs.microsoft.com/en-us/azure/azure-functions/functions-overview)
- [GCP Functions](https://cloud.google.com/functions/)
- etc

---

### Why?

* Time to Market
* Low cost
* Easy integration (API Gateway, Kinesis, SNS, S3)
* NoOPS (almost!!!)
* Deployment (cli)
* Stages (Dev, Staging, Prod)
* Versioning/Rollback
* Automatic Scaling
* Logs build in (at least for AWS)
* Free plans available

---

### [serverless.com framework](https://serverless.com/)

* all in one solution
* powered by CloudFormation but abstracted away complexity
* powerful and easy cli
* multi language support (js, C#, Java, Python)
* multi cloud provider support (AWS, Azure etc)
* easy integration with cloud services (SNS, Kinesis, Api Gateway)
* good documentation

***

### CONS

* High cost if slow process time with large memory need and al ot of requests
* Startup cost is high
* Not everything is suited for this like
    * high performance service
    * process that spans over several minutes
    * ultra-low latency requirements    
    * process that need disk for processing

***

### Why NodeJS? 

* Fast Lambda start up times compared to C#, Java
* Less Resources used compared to C#, Java
* Typescript (thank god)
* npm (package manager like nuget)
* huge package repository
* Small deployment size compared to C#, Java
* Speed
* Easy testing (mocha, chai)

***

### Typescript?
### Is Javascript not bad enough?

* Based on future ECMAScript Specs
* Types and checks at build time (a clear winner)
* transpile future ES6 to ES5
* class, promises, async-await, generics, discriminated unions
* Very familiar to C# (Java, C++) Developers
* [Visual Studio Code](https://code.visualstudio.com/) with autocomplete and debugging support
* Easy to learn
* Created by [Anders Hejlsberg](https://en.wikipedia.org/wiki/Anders_Hejlsberg) (Turbo Pascal, Delphi, C#)

***

### Hands On Goals

* Use `sls` cli to create a service in NodeJS
* Create two functions in the handler
    * retrieve, which returns a list of orders
    * create, which creates a new order
    * test locally with `sls` cli
* Use Api Gateway to create a GET /Orders endpoint
* Deploy using `sls`
* Test in AWS
    * use Postman to send GET/POST
    * monitor logs using `sls`
* Make a change,deploy and test
* Rollback due to a ***bug***

***

### Thank you!

* [presentations](https://github.com/mantzas/presentations)
