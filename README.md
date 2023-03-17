to build the todo-app containers run
```bash
./build-todo-app.sh <version>
```
then you will need to verify the version and push both versions
```bash
docker push gcr.io/solo-test-236622/graphql-todo:0.0.3-amd64
docker push gcr.io/solo-test-236622/graphql-todo:0.0.3-arm64"
```