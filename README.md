# minikube-k8s-demo-app

# Prerequisites

* Docker: Ensure Docker is installed and running on your machine.
* Golang: Install Golang to build the application.
* Minikube: Set up Minikube to create a local Kubernetes cluster.

# Steps to Run the Application Locally

1. Build the Application Image:
    ```
    // build image using docker build which is configured in the Makefile.
    ~ make build-service
    ```
    * This command will build the Docker image for your application, defined in the Makefile.

2. Push the Image to a Registry (Optional):
    ```
        make push-service
    ```
    * If you want to push the image to a registry like Docker Hub, use this command, in makefile it expects a valid repo path.


3. Create a Kubernetes Namespace:
    ```
        kubectl create namespace <namespace-name>
        OR 
        kubectl apply -f ./iac/k8s/application-namespace.yaml -n <namespace>
    ```
    * Replace <namespace-name> with your desired namespace name.

4. Create Configuration and Secret:
    ```
        kubectl apply -f ./iac/k8s/application-configmap.yaml -n <namespace-name>
        
        kubectl apply -f ./iac/k8s/application-secret.yaml -n <namespace-name> 
    ```
    * These commands will create a ConfigMap and Secret in the namespace to store configuration and sensitive information.
5. Deploy the Application:
    ```
        kubectl apply -f ./iac/k8s/application-deployment.yaml -n <namespace-name>
        
        kubectl apply -f ./iac/k8s/application-service.yaml -n <namespace-name>
    ```

    * These commands will deploy the application to the Kubernetes cluster and expose it as a service.

6. Check Application Status:
    ```
        kubectl get pods -n <namespace-name>
    ```
    * This command will list the pods associated with your application.

7. Access the Application:

    * Minikube Dashboard: Use the Minikube dashboard to visualize your cluster and access the application's service.
    * kubectl Port-Forwarding: Use kubectl port-forward to expose the service locally.
    * NodePort Service: If you need external access, configure the service with type: NodePort.

## Additional Notes:

* Mongo Deployment:
  ```
    kubectl apply -f ./iac/k8s/mongo-5.0.yaml -n <namespace-name>
  ```
    * This command deploys a MongoDB instance for the application's data storage needs.
Customizing the Deployment:

* Modify the YAML files in the iac/k8s directory to customize the deployment, such as scaling, resource limits, and environment variables.

* Troubleshooting:

    * Use kubectl logs to view logs from your application's pods.
    * Check the Kubernetes events for error messages.
    * Use kubectl describe pod to get detailed information about a pod.
