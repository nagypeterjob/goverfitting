#!/bin/bash
function getContainersClean() {
  echo "Wiping out stopped containers"
	for i in $(docker ps -a -q);
	do
		docker rm $i &>/dev/null
	done
}

function getImagesClean() {
  echo "Wiping out none tagged images"
	for i in  $(docker images -a | grep "^<none>" | awk '{print $3}')
	do
		docker rmi $i &>/dev/null
	done
}

getContainersClean
getImagesClean
exit