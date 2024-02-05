# First Custom Kubernetes Controller

Welcome to the repository for my first custom Kubernetes controller! This project is a deep dive into extending Kubernetes' functionality through the development of a custom controller. It serves as a practical exploration of Kubernetes internals, Go programming, and the power of custom resources within the Kubernetes ecosystem.

## Overview

This custom Kubernetes controller is designed to manage a specific type of custom resource, demonstrating how to watch, create, update, and delete Kubernetes resources programmatically. The controller uses `client-go`, the official Go client library for Kubernetes, to interact with the Kubernetes API.

## Features

- Custom Resource Definition (CRD) setup and management
- Watching for changes to specific resources in the Kubernetes cluster
- Reconciliation logic to ensure the desired state of resources
- Example custom resources for testing the controller


