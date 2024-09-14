# Sample application for GraphQL

## Features

* *Multi stage*: Uses Multi stage build for building the container
* *Database ownership*: Takes care of the database schema by itself
* *Configuration*: all runtime parameters are passed either via environment variables or command line arguments
  * Kubernetes ConfigMap
  * Kubernetes Secrets
* *Logging*:
  * JSON logging
  * Logging of trace_id and span_id for correlation
* Authentication: Validate the JWT token
  * Token check can be disabled for easier development
  * The following entries are written to the span
    * enduser.id
    * enduser.roles
* OpenTelemetry: Exposes traces in OTLP to a configurable endpoint
* *Probes*: Exposes readiness and liveness probes
  * /health endpoint
* Version awareness: Exploses the own version via /version endpoint
* Automated testing

## Diagram

![graph.png](graph.png)

# Queries

## TagCategories and Tags

```
query {
  tagCategories {
    id
    name
    __typename
    parentTagCategory {
      id
      name
      __typename
    }
    ...on StaticTagCategory {
      isOpen
    }
    ...on DynamicTagCategory {
      format
    }
    childTagCategories {
      id
      name
      __typename
    }
    rootTags {
      id
      ...on StaticTag {
        name
      }
      ...on DynamicTag {
        value
      }
      parentTag {
        id
        ...on StaticTag {
          name
        }
        ...on DynamicTag {
          value
        }
      }
      childTags {
        id
        ...on StaticTag {
          name
        }
        ...on DynamicTag {
          value
        }
      }
    }
  }
}
```

## Asset with Tags

```
query AssetsWithTags {
  asset {
    id
    name
    files {
      id
      name
      size
      mimeType
    }
    tags {
      __typename
      id
      ...on StaticTag {
        name
        tagCategory {
          id
          name
          parentTagCategory {
            id
            name
          }
          isOpen
        }
      }
      ...on DynamicTag {
        value
        tagCategory {
          id
          name
        }
      }
    }
  }
}
```