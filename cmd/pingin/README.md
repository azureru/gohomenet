# Pingin (Ping Interface)

So my setup is one Pi/OPi/SBC computer with multiple network interfaces that are
connected to multiple ISP (for redundancy :P)

So this utility basically will send HTTP GET request to each specified interfaces
on PINGIN_INTERFACE env and then give appropriate output for each interface (basically `curl --interface`)

On my case the HTTP endpoint is basically a webhook to record data to influx+grafana setup - so as a bonus I graph the connectivity stability of
each ISP 
