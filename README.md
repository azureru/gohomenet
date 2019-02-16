# gohomenet
Notes on experimentation with network at my home using golang

## The Hardware
- A orangepi with one Ethernet and one USB-to-Ethernet
- One link to ISP I*
- One link to ISP F*

# Multiple Interface Routing
To setup the multi routing for the above setup I follow guide
https://www.thomas-krenn.com/en/wiki/Two_Default_Gateways_on_One_System
After that we can test the network using
```
curl --interface eth0 http://whatismyip.akamai.com/
curl --interface eth3 http://whatismyip.akamai.com/
```
It should show different public IP.

#
