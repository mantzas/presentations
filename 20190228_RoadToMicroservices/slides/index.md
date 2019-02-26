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

## Avoid HDD

Hype Driven Development

---

> don’t even consider microservices unless you have a system that’s too complex to manage
> as a monolith. Martin Fowler

---

## They are no silver bullets

***

## Distributed Systems are hard

---

### [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)

- The network is reliable.
- Latency is zero.
- Bandwidth is infinite.
- The network is secure.
- Topology doesn't change.
- There is one administrator.
- Transport cost is zero.
- The network is homogeneous.

---

### [CAP Theorem](https://en.wikipedia.org/wiki/CAP_theorem)

> ... due to the fact that network partitioning has to be tolerated 
> there are only Consistency or Availability.

---

### [PACELC theorem (CAP theorem extension)](https://en.wikipedia.org/wiki/PACELC_theorem)

> ... in case of network partitioning (P) in a distributed computer system, 
> one has to choose between availability (A) and consistency (C) (as per the CAP theorem), 
> but else (E), even when the system is running normally in the absence of partitions, 
> one has to choose between latency (L) and consistency (C).

***

## Problem: Failures

- Errors
- Slow response times

---

## Solution: Failures

- [Circuit Breaker](https://martinfowler.com/bliki/CircuitBreaker.html)
- Proper Timeouts
- Fail fast

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

## Problem: Database

- Which one?
- More than one?

---

## Solution: Database

- Any, not confined by monolith's choice
- You should be able to support it though...

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

***

## Problem: How to start

- [Coupling](https://en.wikipedia.org/wiki/Coupling_(computer_programming))
- Distributed Monolith

## Solution: How to start

- [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
