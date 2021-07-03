#! /bin/sh
sudo cpufreq-set -g performance

index=0
index1=0
#mount hard disk
for V in $(ls /dev/sd[b-z][1-26])
do
 echo $V

 sudo ntfsfix $V

M=$(echo "$V" |awk -F "/" '{print $3}')
#echo $M
mkdir -p /home/cykj/HHD${index} &
sudo chmod -R 777 /home/cykj/HHD${index}
sudo mount -t ntfs $V /home/cykj/HHD${index} &
index=`expr $index + 1`
done

#mount hard disk without number and it's start with 'c'
for V in $(ls /dev/sd[c-z])
do
 echo $V

# sudo ntfsfix $V

M=$(echo "$V" |awk -F "/" '{print $3}')
#echo $M
mkdir -p /home/cykj/HHD${index} &
sudo chmod -R 777 /home/cykj/HHD${index}
sudo mount -t ext4 $V /home/cykj/HHD${index} &
index=`expr $index + 1`
done

# format ssd
for V in $(ls /dev/nvme[0-9]n1p[0-9])
do
	str=${V:0:12}
 echo $str
sudo fdisk $str <<EOF
d
2
d
d
w
EOF
sudo mkfs.ext4 $str <<EOF
y
EOF
done

#:<<EFO
#mount sand disk
for V in $(ls /dev/nvme[0-9]n1)
do
 echo $V
M=$(echo "$V" |awk -F "/" '{print $3}')
#echo $M
mkdir -p /home/cykj/SSD${index1} &
sudo chmod -R 777 /home/cykj/SSD${index1}
sudo mount -t ext4 $V /home/cykj/SSD${index1} &
index1=`expr $index1 + 1`
done
#EFO
