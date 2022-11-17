
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
docker buildx build --platform $ARCH -t $REPO/users:$VERSION --no-cache  --push .

cd $PARENT_PWD
cd ./users/users-rest 
docker buildx build --platform $ARCH -t $REPO/users-rest:$VERSION --no-cache  --push .


#Comments
cd $PARENT_PWD
cd ./comments
docker buildx build --platform $ARCH -t $REPO/comments:$VERSION --no-cache  --push .

docker buildx rm
