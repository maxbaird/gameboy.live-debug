#! /bin/bash

i=$1

while [ $i -le $2 ]
do
    #echo "cmp /tmp/$i.gblive /tmp/$i.topaz"
    cmp /tmp/$i.gblive /tmp/$i.topaz
    if [[ $? -ne 0 ]]
    then
     echo "Comparsion failed at state $i"
    fi
    ((i++))
done

echo "Done"
