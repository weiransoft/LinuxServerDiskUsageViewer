LinuxServerDiskUsageViewer
==========================

LinuxServerDiskUsageViewer is distributed under a BSD-style license.

This viewer is developed by GoLang.
===================================================
用于查看Linux 下的磁盘空间查看

查看整个挂载磁盘的使用情况
http://ip:8080/viewDiskUsage

查看特定目录的空间占用情况
加上/view?path= 目录地址
如:
http://ip:8080/viewDiskUsage/view?path=/opt/mount/

===================================================
It's use to view the Linux server disk usage.

Monitor whole mounted disk usage:

http://ip:8080/viewDiskUsage

Monitor path which you want:

Ex:
http://ip:8080/viewDiskUsage/view?path=/opt/mount/

Getting Started
0.Setup your GoLang development environment
1.Clone this project to your eclipse or Linux.
2.Execute Go command:Go install LinuxServerDiskUsageViewer
3.Use LinuxServerDiskUsageViewer 
4.Access http://ip:8080/viewDiskUsage
