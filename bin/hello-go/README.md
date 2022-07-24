# Hello Go

## Workflow

```bash
# Build image
docker build -t maroskukan/hello-go .

# Run container
docker run --rm -p 8180:8180 maroskukan/hello-go

# Test container
curl localhost:8180

# Push image
docker push maroskukan/hello-go
```