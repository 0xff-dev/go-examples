#!/bin/bash
echo "tag images\nfirst please login registroy"
for key in mongo mc
do
    docker images | grep $key | cat |  while read line 
    do
        tag=`echo $line | awk '{print $3}'`
        newImage=`echo $line | awk '{newImage=$1":"$2;print newImage}' | sed 's/192.168.0.58/192.168.1.52/g'`
        echo "run :=====> docker tag $tag $newImage"
        #docker rmi $newImage
        docker tag $tag $newImage
        docker push $newImage
    done
done
