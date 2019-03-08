# ixstorage Eventsourcing Quick Start Guide

The communication between ixstorage services is initiated through event sourcing.

## What is Event Sourcing?

## Why is Event Sourcing useful?

For communicating between ixstorage services we decided to implement an event sourced system which enables [CQRS](https://martinfowler.com/bliki/CQRS.html) and with that provides a good amount of decoupling between our services. Also, we want to ensure that each event occuring on our aggregates can be reproduced at a later point in time if necessary.