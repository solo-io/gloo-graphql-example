
## Builder
docker buildx create --name mybuilder
docker buildx use mybuilder



REPO="${REPO:-asayah}"
VERSION="${VERSION:-latest}"
ARCH="linux/amd64,linux/arm64"

PARENT_PWD=$PWD

#Blogs
cd ./blogs
docker buildx build --platform $ARCH -t $REPO/blogs:$VERSION --no-cache  --push .

#Users 
cd $PARENT_PWD
cd ./users/users-apollo 
docker build --platform linux/amd64 -t $REPO/users-apollo:$VERSION --no-cache . 
docker push $REPO/users-apollo:$VERSION

cd $PARENT_PWD
cd ./users/users-rest 
docker build --platform linux/amd64 -t $REPO/users-rest:$VERSION --no-cache . 
docker push $REPO/users-rest:$VERSION

#Comments
cd $PARENT_PWD
cd ./comments
docker build --platform linux/amd64 -t $REPO/comments:$VERSION --no-cache . 
docker push $REPO/comments:$VERSION

