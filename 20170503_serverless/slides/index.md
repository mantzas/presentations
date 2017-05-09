- title : serverless
- description : Introduction into serverless 
- author : Sotirios Mantziaris
- theme : night
- transition : default

***

## serverless

<br />
<br />

### serverless intro with Node.js and Typescript on AWS Lambda

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

### What about speed of provisioning?

#### On Premises

> Bare Metal ~ ***weeks?***
<br />
> VM ~ ***hours? days? weeks?***
<br />
> Docker ~ ***minutes? days? weeks?***

#### Cloud 

> VM ~ ***hours? days?***
<br />
> Docker ~ ***minutes? days?***
<br />
> serverless ~ ***seconds?***

***

### What is serverless?

> Serverless computing, also known as function as a service (FaaS), is a cloud computing code execution model in which the cloud provider fully manages starting and stopping of  a function's container platform as a service (PaaS) as necessary to serve requests, and requests are billed by an abstract measure of the resources required to satisfy the      request, rather than per virtual machine, per hour.
 [Wikipedia](https://en.wikipedia.org/wiki/Serverless_computing)

---

### What does that actually mean?

- No provisioning of Infrastructure
- Pay by usage not by hour (used or not)

---

### Who offers serverless

- [AWS Lambda](https://aws.amazon.com/lambda/)
- [Azure Functions](https://docs.microsoft.com/en-us/azure/azure-functions/functions-overview)
- [GCP Functions](https://cloud.google.com/functions/)
- etc
<br />
<br />
***I will concentrate on AWS for the rest of the presentation***

---

### Why?

* Time to Market
* Low cost
* Easy integration (API Gateway, Kinesis, SNS, S3)
* NoOPS (almost!!!)
* Easy deployment (cli)
* Stages (Dev, Staging, Prod)
* Versioning/Rollback
* Automatic Scaling
* Logs build in (at least for AWS)
* Free plans available

---

### CONS

* High cost if slow process time with large memory need and a lot of requests
* Function startup cost is high
* Function lives only for 300 sec (AWS, Azure etc)
* Not everything is suited for this
    * high performance service
    * process that spans over several minutes
    * ultra-low latency requirements    
    * process that need disk for processing

***

### I like it, how do i start?

* Via the provider website (upload code etc)
* Via [Visual Studio Extensions](https://aws.amazon.com/blogs/developer/preview-of-the-aws-toolkit-for-visual-studio-2017/) for C#
* Via [dotnet cli](http://docs.aws.amazon.com/toolkit-for-visual-studio/latest/user-guide/lambda-cli-publish.html) for C#
* Via [serverless framework](https://serverless.com/)
<br />
<br />
***For the rest of the presentation i will use the serverless framework which makes the process very easy***

---

### [serverless.com framework](https://serverless.com/)

* tries to be a all in one solution
* powered by CloudFormation, but abstracted away complexity of it
* powerful and easy cli
* multi language support (js, C#, Java, Python)
* multi cloud provider support (AWS, Azure etc)
* easy integration with cloud services (SNS, Kinesis, Api Gateway)
* good documentation
* **you may have to dig deeper for highly advanced stuff**

***

### What's the stack?

* [serverless framework](https://serverless.com/)
* [Node.js](https://nodejs.org/en/)
* [Typescript](https://www.typescriptlang.org/)

---

### Why NodeJS? 

* Faster Lambda start up times compared to C#, Java
* Less Resources used compared to C#, Java
* [Typescript](https://www.typescriptlang.org/) (thank god)
* [npm](https://www.npmjs.com/) (package manager like nuget)
* huge package repository
* Smaller deployment size compared to C#, Java
* Speed
* Easy testing ([mocha](https://mochajs.org/), [chai](http://chaijs.com/))

---

### Typescript?
### Is Javascript not bad enough?

* Based on future ECMAScript Specs
* Types and checks at build time (a clear winner)
* transpile future ES6 to ES5
* classes, promises (like Task), async-await, generics, discriminated unions
* Very familiar to C# (Java, C++) Developers
* [Visual Studio Code](https://code.visualstudio.com/) with autocomplete and debugging support
* Easy to learn
* Created by [Anders Hejlsberg](https://en.wikipedia.org/wiki/Anders_Hejlsberg) (Turbo Pascal, Delphi, C#)

***

### Enough with the theory
# ***Show me the code***

---

### First we pray to the demo gods!!!

![Demo](./images/Demo.jpeg)

---

### Hands On Goals

* Use `sls` cli to create a service in NodeJS
* Create a function in the handler
    * retrieve, which returns a list of orders
    * test locally with `sls` cli
* Use Api Gateway to create a GET /Orders endpoint
* Deploy using `sls`
* Test in AWS
    * use Chrome to send GET
    * monitor logs using `sls`
* Make a change,deploy and test
* Rollback due to a ***bug***
* Remove the service

***

### Thank you!

* [presentations](https://github.com/mantzas/presentations)
