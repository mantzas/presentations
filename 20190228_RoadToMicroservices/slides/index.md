- title : Road to Microservices
- description : Introduction to FsReveal
- author : Sotirios Mantziaris
- theme : night
- transition : default

***

# Road to Microservices

Sotirios Mantziaris  
Refreshment Engineer (καφετζής)

smantziaris@gmail.com  
<http://blog.mantziaris.eu>  
<https://github.com/mantzas>  
<https://gr.linkedin.com/in/mantzas>  
@smantziaris

***

## From Monolith to microservices

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

## Problem and Solutions

***

## Problem: Failures

- Error handling
- Slow response times

---

## Solution: Failures

- [Circuit Breaker](https://martinfowler.com/bliki/CircuitBreaker.html)
- Proper Timeouts
- Generally fail fast

***

## Problem: Infrastructure

- Orchestrate microservices
- Continuous deployments

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

## Problem: Configuration & Secrets

- Configs
- Secrets

---

## Solution: Configuration & Secrets

- [Consul](https://www.hashicorp.com/products/consul)
- [Vault](https://www.hashicorp.com/products/vault/)

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
- statically typed, statically linked, one executable
- Concurrent
- Fast and Furious
- Easy to learn
- Easy to maintain
- Great tooling (testing, profiling etc.)
- Comprehensive std packages and rich 3rd party packages
- Docker friendly e.g. `FROM SCRATCH` creates executable sized container
- see you @ [Athens Gophers Meetup](https://www.meetup.com/Athens-Gophers/)

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

- Integrated all the above best practices
- Support for sync component e.g. HTTP
- Support for async component e.g. Kafka, AMQP (RabbitMQ)
- Support other components via simple interface
- Component handler are agnostic of the underlying delivery
- Observability: Structured logging, Prometheus metrics, Jaeger distributed tracing
- Clients with integrated trace propagation e.g. HTTP, Kafka producer, AMQP publisher
- Open Source, Apache 2 License
- [Beat](https://github.com/thebeatapp/patron) forked and adopted

***

## Problem: How to start

- [Coupling](https://en.wikipedia.org/wiki/Coupling_(computer_programming))
- Distributed Monolith

---

## Solution: How to start

- [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
- Start small with a good defined domain
- [Strangler Pattern](https://developer.ibm.com/articles/cl-strangler-application-pattern-microservices-apps-trs/)

***

## Enjoy the ride

***

## Q&A

***

## Stuff

- Slides - <https://github.com/mantzas/presentations>
- Patron - <https://github.com/mantzas/patron>
- Blog - <https://blog.mantziaris.eu>
- Athens Gophers Meetup - <https://www.meetup.com/Athens-Gophers>