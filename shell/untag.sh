for key in "$*"
do
    docker images | grep $key | cat | while read line
    do
       tag=`echo $line | awk '{print $1":"$2}'`
       echo "docker rmi $tag"
       docker rmi $tag
    done
done
