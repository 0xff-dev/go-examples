images=("quay.io/cephcsi/cephcsi:v3.1.1" "quay.io/k8scsi/csi-node-driver-registrar:v1.2.0" "quay.io/k8scsi/csi-resizer:v0.4.0" "quay.io/k8scsi/csi-provisioner:v1.6.0" "quay.io/k8scsi/csi-snapshotter:v2.1.1" "quay.io/k8scsi/csi-attacher:v2.1.0" "rook/ceph:v1.4.8" "ceph/ceph:v14.2.10")
harborAddrss="192.168.1.52"
project="system_containers"
pullImage=0

while getopts 'dH:P:' arg
do
    case $arg in
        d)
            pullImage=1
            ;;
        H)
            harborAddrss=$OPTARG
            ;;
        P)
            project=$OPTARG
            ;;
    esac
done


for name in ${images[@]}
do
    newTag=$harborAddrss/$project/$(echo $name | awk -F/ '{print $NF}')
    echo "image: $newTag"
    if [ $pullImage -eq 1 ];
    then
        echo "docker pull $name"
        # docker pull $name
    fi
    echo "docker tag $name $newTag"
    #docker tag $name $harborAddrss$(echo $name | awk -F/ '{print $NF}')
    echo "docker push $newTag\n"
done

