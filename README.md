# Health Service Provider

A quick POC to demonstrate how to create a service provider for Gostalt.

This creates a "health" service provider that has a list of checks (e.g., CPU,
number of 500s).

These checks will be fired as part of the schedule, and we also surface a route
that can contain the status of these checks.
