- title : Road to Microservices
- description : Introduction to FsReveal
- author : Sotirios Mantziaris
- theme : night
- transition : default

***

## Road to Microservices

### From Monolith to microservices

Sotirios Mantziaris  
Refreshment Engineer (καφετζής)

smantziaris@gmail.com  
http://blog.mantziaris.eu  
https://github.com/mantzas  
https://gr.linkedin.com/in/mantzas  
@smantziaris  

***

## They are no silver bullets

- Avoid HDD (Hype Driven Development)

- Another good point made by Martin Fowler is:

> don’t even consider microservices unless you have a system that’s too complex to manage
> as a monolith.

***

## Distributed Systems are hard

***

### [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)

- The network is reliable.
- Latency is zero.
- Bandwidth is infinite.
- The network is secure.
- Topology doesn't change.
- There is one administrator.
- Transport cost is zero.
- The network is homogeneous.

***

### [CAP Theorem](https://en.wikipedia.org/wiki/CAP_theorem)

> ... due to the fact that network partitioning has to be tolerated 
> there are only Consistency or Availability.

***

### [PACELC theorem (CAP theorem extension)](https://en.wikipedia.org/wiki/PACELC_theorem)

> ... in case of network partitioning (P) in a distributed computer system, 
> one has to choose between availability (A) and consistency (C) (as per the CAP theorem), 
> but else (E), even when the system is running normally in the absence of partitions, 
> one has to choose between latency (L) and consistency (C).

***

## Problem: Infrastructure

- Orchestrate multiple services
- Manage Deployments

---

## Solution: Infrastructure

- Orchestration: [Kubernetes](https://kubernetes.io/)
- Deployments: [Docker Containers](https://www.docker.com/why-docker), [Helm](https://helm.sh/) and CI/CD of your choice

***

## Problem: Observability

- Logging
- Metrics
- Distributed Tracing

---

## Solution: Observability

- Logging: [ELK](https://www.elastic.co/elk-stack)
- Metrics: [Prometheus](https://prometheus.io/) (or [InfluxDB](https://www.influxdata.com/)) and [Grafana](https://grafana.com/)
- Distributed Tracing: [Opentracing](https://opentracing.io/): [Jaeger](https://www.jaegertracing.io/)

***

## Problem: Programming Language

- Which One?
- More than one?

---

## Solution: Language

- Any language that is easily containerized
- Adopt 2 or 3

> *but look especially into...*

---

## Solution: [Go](https://golang.org/)

- *Cloud native*
- Concurrent
- Fast
- Easy to learn
- Maintainable
- Great tooling (testing, profiling etc.)
- Comprehensive std packages and 3rd party packages
- Docker friendly (FROM SCRATCH)

> Almost every provided solution here is written in go e.g. Kubernetes, Docker, Prometheus etc.

***

## Problem: Any help

- Framework?

---

## Solution: Any help

- [Spring Cloud Java](https://spring.io/projects/spring-cloud)
- [Flask](http://flask.pocoo.org/)
- [Go kit](https://gokit.io/)

> *but look especially into...*

---

## Solution: [Patron](https://github.com/mantzas/patron)

- 
- [Beat](https://github.com/thebeatapp/patron) forked and adopted

## How to carve out those things?

- DDD

***

* Response times (RT) increases (optimistic)

user --> Service A --> Service B --> Service C

where:

    RT Service A: 60ms
    RT Service B: 30ms
    RT Service C: 50ms

Total response time for the user is:

    RT Service C + RT Service B + RT Service A = 140ms

which will be generally bigger than the response time of a monolith due to network, marshal/unmarshal data etc.

* Response times (RT) explodes (pessimistic)

user --> Service A --> Service B --> Service C

where:

    RT Service A: 60ms
    RT Service B: 30ms
    RT Service C: times out at 60s

Total response time for the user is:

    RT Service C + RT Service B + RT Service A = 1m 90ms

which would not happen in a monolith at all!

* Availability decreases

e.g.

    Availability Service A: 99.9%
    Availability Service B: 99.9%
    Availability Service C: 99.9%

the joint availability ([[https://en.wikipedia.org/wiki/Probability#Mathematical_treatment][probability rules]], independent events) for the user is:

    Availability Service A * Availability Service B * Availability Service C ~ 99.7%

Even if all services are 99.9% the joint availability is lower than the individual ones.

The difference is not that small! ([[https://uptime.is/][downtime calculation]])

    99.9 ~ 8h 45m 57.0s downtime
    99.7 ~ 1d 2h 17m 50.9s downtime

* Coupling

[[https://en.wikipedia.org/wiki/Coupling_(computer_programming)][coupling]] is the degree of interdependence between software modules.

how can this happen?

- microservice creation based on opportunity
- shared db
- access the db directly and not through a API
- shared components, like models and business code

the above will lead to what is commonly called a "distributed monolith".

* Indications of a "distributed monolith"?

- A change to one microservice often requires changes to other microservices
- Deploying one microservice requires other microservices to be deployed at the same time
- Your microservices are overly chatty
- The same developers work across a large number of microservices
- Many of your microservices share a data storage
- Your microservices share a lot of the same code or models

* Troubleshooting failures

In order to troubleshoot you have to look into multiple services to find the problem.

- Metrics
    help you identify the general area of a problem e.g. response time is high

- Distributed tracing
    can narrow down the problem e.g. response time of the db is high on service A

- Logging
    can show details tha are logged eg what happened to payment with id xxx

all the above fall under the category of *observability*.

* how can any or all of the previous be avoided?

- DDD can help, e.g. separate bounded contexts
- Observability with logging, metrics, distributed tracing
- Circuit breakers, correct timeouts in place
- Not everything has to be microservices
- Data duplication can be helpful


***

**Bayes' Rule in LaTeX**

$ \Pr(A|B)=\frac{\Pr(B|A)\Pr(A)}{\Pr(B|A)\Pr(A)+\Pr(B|\neg A)\Pr(\neg A)} $

***

### The Reality of a Developer's Life 

**When I show my boss that I've fixed a bug:**
  
![When I show my boss that I've fixed a bug](http://www.topito.com/wp-content/uploads/2013/01/code-07.gif)
  
**When your regular expression returns what you expect:**
  
![When your regular expression returns what you expect](http://www.topito.com/wp-content/uploads/2013/01/code-03.gif)
  
*from [The Reality of a Developer's Life - in GIFs, Of Course](http://server.dzone.com/articles/reality-developers-life-gifs)*

