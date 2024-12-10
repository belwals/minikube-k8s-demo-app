#!/bin/bash

NAMESPACE=$1
if [$NAMESPACE == ""] 
then 
    echo "Namespace is not provided, using application as default namespace"
    NAMESPACE=application
fi

kubectl delete namespace $NAMESPACE
