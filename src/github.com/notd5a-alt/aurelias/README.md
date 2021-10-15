# Go-NMAP
    A simple network scanner made using Go


## Objective: 

Given a network address you need to scan all the IPs to prepare a list which all are up. Then scan the IPs which are up for the well known ports to see if they are accepting connections. Scan telnet, ftp, http, ssh etc. At the end print a full report of the network.

### Ideas
1. Allow for multiple types of scans:
    1. TCP (Lower range scan)
    2. UDP (Higher range scan)
    3. Wide Range Scans (searching both)
    4. Ping Scans (Quick)
2. Cache the data from the previous scan
3. Allow for an option to store the data of the scan to a json file

