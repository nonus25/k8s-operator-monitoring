# k8s-operator-monitoring
- health Prometheus endpoint
- Health of different components + metrics
- Operator can depend on this Endpoint to monitor the status and try to correct when necessary.

## 
 1. Explore how to expose an health endpoint in the Operator
 2. Define the metrics which this endpoint checks and relays
 3. Operator itself will keep monitoring this endpoint
     - if any component is unhealthy, the Operator reacts with operations like (restart pod, restart the deployment etc..) ???
