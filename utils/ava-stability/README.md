# Stress and Stability Test

## Abstract
Stress and stability testing was conducted on AvalancheGo 1.3.2. It was observed during many performance tests that with multi-node setups split between and HTTP API tier and a consensus tier, TPS would top out while resource utilization on the nodes (like CPU, disk, network, etc.) would remain relatively low. The only way found to increase TPS would be to add more API nodes, and it did not appear adding more CPU's per API node made any difference in throughput. Only more physical nodes made a difference. The goal was to investigate the source of the bottleneck causing lack of resource utilization to see if TPS per node could be increased.

## Setup
It was decided to simplify the hardware and software test setup as much as possible. This would reduce investigative complexity. The only problem with this approach is that it would not fully exercise consensus, and thus may exhibit different behavior than multi-node setups. But it was thought that if that were the case, the data would bear that out, and thus provide insight into the next investigative direction to pursue.

## Client and Server Testing Hardware
- One single Macbook Pro
  - 2019 edition
  - 8 core, dual thread 2.4 Ghz Core i9
  - 64 GB 2667 MHz DDR4 RAM
  - 2 TB APPLE SSD AP2048N
  - OSX 10.15.7

## Client Software
A [stress test application][1] was written in javascript and executed with node.js, V16.4.2. This application contains [one smart contract][2] with a simple method to increment a instance variable. The validator scdev chain contains over 1000 wallet addresses with plenty of Flare, and so one smart contract per wallet address was deployed. The node application then issued one promise per wallet/contract to increment the instance variable. Promise.race was then used to catch promise completion and re-issue another promise and this was put into am endless loop. The application was thus fully asynchronous pushing all load/waiting activity to the validator application.

The command line of the stress application takes one parameter `-t` specifying the number of contracts to deploy and `-o` specifying the offset in the array of addresses to begin referencing the first wallet address. The offset enables multiple client applications to be run simultaneously without worrying about nonce collisions.

The client software counts the number of successful increment calls made to each respective smart contract, measures the time elapsed for the transaction to complete and fetch back the counter. TPS is a measure of the number of these transactions completed over the time of total application execution. Average elapsed time, max elapsed time, and min elapsed time are also tracked. These values are reported to the console every 5 seconds. When a transaction completes, it's contract index number, address, and count are also reported to the console.

## Server Software
Flare Networks customized version of Ava Lab's AvalancheGo application was deployed. The commit hash of the branch created was  `e9ca17eace0960e34b67cbb08acb43dc9092f613`. Note that in the interim, Flare's version was upgraded to AvalancheGo 1.4.10. This test did not take those updates.

AvalancheGo was deployed on the hardware mentioned, in a two tiered deployment with a "1-node" API layer and a "1-node" consensus layer. In fact, both instances were deployed on the same physical node for convenience. The deployment script follows:

```bash
#!/bin/bash
# If you pass --existing, then the script will bypass the build.
# If you pass --clean, any existing chain database will be blasted.
if [[ $(pwd) =~ " " ]]; then echo "Working directory path contains a folder with a space in its name, please remove all spaces" && exit; fi
if [ -z ${GOPATH+x} ]; then echo "GOPATH is not set, visit https://github.com/golang/go/wiki/SettingGOPATH" && exit; fi
if [ -z ${XRP_APIs+x} ] || [ "$XRP_APIs" == "url1, url2, ..., urlN" ]; then echo "XRP_APIs is not set, please set it using the form: $ export XRP_APIs=\"url1, url2, ..., urlN\"" && exit; fi
if [[ $(go version) != *"go1.15.5"* ]]; then echo "Go version is not go1.15.5" && exit; fi
XRP_APIs_JOINED="$(echo -e "${XRP_APIs}" | tr -d '[:space:]')"
printf "\x1b[34mFlare Network 1-Node Smart Contract Dev Team Static IP Address Dual Node Local Deployment\x1b[0m\n\n"
AVALANCHEGO_VERSION=@v1.3.2
CORETH_VERSION=@v0.4.2-rc.4

EXEC_DIR=$(pwd)
LOG_DIR=$(pwd)/logs
CONFIG_DIR=$(pwd)/config
PKG_DIR=$GOPATH/pkg/mod/github.com/ava-labs
NODE_DIR=$PKG_DIR/avalanchego$AVALANCHEGO_VERSION
CORETH_DIR=$PKG_DIR/coreth$CORETH_VERSION

if echo $1 | grep -e "--existing" -q
then
	cd $NODE_DIR
  if echo $2 | grep -e "--clean" -q
  then
    echo "Cleaning chain db..."
    rm -rf $NODE_DIR/db/
  fi
else
	rm -rf logs
	mkdir logs
	rm -rf $NODE_DIR
	mkdir -p $PKG_DIR
	cp -r fba-avalanche/avalanchego $NODE_DIR
	rm -rf $CORETH_DIR
	cp -r fba-avalanche/coreth $CORETH_DIR
	cd $NODE_DIR
	rm -rf $NODE_DIR/db/
	mkdir -p $LOG_DIR/node00
	mkdir -p $LOG_DIR/node01
	printf "Building Flare Core...\n"
	./scripts/build.sh
fi

# NODE 0
printf "Launching Node 0 at 192.168.2.1:9660\n"
./build/avalanchego \
    --public-ip=192.168.2.1 \
    --snow-sample-size=1 \
    --snow-quorum-size=1 \
    --http-port=9660 \
    --staking-port=9661 \
    --db-dir=$(pwd)/db/node00/ \
    --staking-enabled=true \
    --network-id=scdev \
    --bootstrap-ips= \
    --bootstrap-ids= \
    --staking-tls-cert-file=$(pwd)/config/keys/node00/node.crt \
    --staking-tls-key-file=$(pwd)/config/keys/node00/node.key \
    --log-level=debug \
    --log-dir=$LOG_DIR/node00 \
    --validators-file=$(pwd)/config/validators/scdev/1619370000.json \
    --alert-apis="https://flare.network" \
    --xrp-apis=$XRP_APIs_JOINED \
    &
sleep 2
NODE_00_PID=`lsof -n -i4TCP:9660 | grep LISTEN | cut -d ' ' -f2`

# NODE 1
printf "Launching Node 1 at 192.168.2.2:9662\n"
./build/avalanchego \
    --public-ip=192.168.2.2 \
    --snow-sample-size=1 \
    --snow-quorum-size=1 \
    --http-host=192.168.2.2 \
    --http-port=9662 \
    --staking-port=9663 \
    --db-dir=$(pwd)/db/node01/ \
    --staking-enabled=true \
    --network-id=scdev \
    --bootstrap-ips=192.168.2.1:9661 \
    --bootstrap-ids=NodeID-D4jrXzkioNZkqbPNuPmk3hR9Ee8oXLvDJ \
    --staking-tls-cert-file=$(pwd)/config/keys/node01/node.crt \
    --staking-tls-key-file=$(pwd)/config/keys/node01/node.key \
    --log-level=debug \
    --log-dir=$LOG_DIR/node01 \
    --validators-file=$(pwd)/config/validators/scdev/1619370000.json \
    --alert-apis="https://flare.network" \
    --xrp-apis=$XRP_APIs_JOINED \
    &
sleep 2
NODE_01_PID=`lsof -n -i4TCP:9662 | grep LISTEN | cut -d ' ' -f2`

printf "\n"
read -p "Press enter to stop background node processes"
kill $NODE_00_PID
kill $NODE_01_PID
```

## Initial Runs
Initial runs of 10, 40, 100, and 200 contracts were performed. It was noted that TPS increased proportionally between 10 and 40 contracts, less than proportionally between 10 and 100 contracts. Runs much over 100 contracts would not result an increase in TPS throughput and runs of 200 contracts and above would result in sporadic TCP connection reset errors on the client and a slight decline in TPS throughput. It was also noted that there would be periods of quiescence in resource utilization and pauses of transaction execution as noted by observing the client console (which prints the result of each execution).

## Analysis
The TCP connection resets where the first clue, pointing to a potential source for trouble analysis. Since it is hard to diagnose any particular transaction conversation from code in a multi-promise setup such as this, it was first thought that since the error messages manifest as networking problems (even though there is no network involved, as everything was running on localhost at that time), that a network analysis tool could be used to spy on networking conversations. For this task, [Wireshark][3] was employed.

The network conversation for an operation in error is easy to spot in Wireshark, as operations in error will show up in red. From an operation in error, a filter can be used to filter only the operations that apply to that specific conversation. An example conversation in error is shown below:

```
No.	Time	Source	Destination	Protocol	Length	Info
58968	16.609381	127.0.0.1	127.0.0.1	TCP	68	50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044080402 TSecr=0 SACK_PERM=1
58975	16.609468	127.0.0.1	127.0.0.1	TCP	68	9660 → 50995 [SYN, ACK] Seq=0 Ack=1 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044080402 TSecr=1044080402 SACK_PERM=1
58983	16.609530	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [ACK] Seq=1 Ack=1 Win=408256 Len=0 TSval=1044080402 TSecr=1044080402
58990	16.609619	127.0.0.1	127.0.0.1	TCP	56	[TCP Window Update] 9660 → 50995 [ACK] Seq=1 Ack=1 Win=408256 Len=0 TSval=1044080403 TSecr=1044080402
59362	16.617774	127.0.0.1	127.0.0.1	HTTP/JSON	437	POST /ext/bc/C/rpc HTTP/1.1 , JavaScript Object Notation (application/json)
59363	16.617789	127.0.0.1	127.0.0.1	TCP	56	9660 → 50995 [ACK] Seq=1 Ack=382 Win=407872 Len=0 TSval=1044080410 TSecr=1044080410
59687	16.715860	127.0.0.1	127.0.0.1	HTTP/JSON	1896	HTTP/1.1 200 OK , JavaScript Object Notation (application/json)
59702	16.715887	127.0.0.1	127.0.0.1	TCP	56	9660 → 50995 [FIN, ACK] Seq=1841 Ack=382 Win=407872 Len=0 TSval=1044080504 TSecr=1044080410
59709	16.715904	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [ACK] Seq=382 Ack=1841 Win=406400 Len=0 TSval=1044080504 TSecr=1044080504
59730	16.715954	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [ACK] Seq=382 Ack=1842 Win=406400 Len=0 TSval=1044080504 TSecr=1044080504
60142	16.742859	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [FIN, ACK] Seq=382 Ack=1842 Win=406400 Len=0 TSval=1044080528 TSecr=1044080504
60143	16.742866	127.0.0.1	127.0.0.1	TCP	56	9660 → 50995 [ACK] Seq=1842 Ack=383 Win=407872 Len=0 TSval=1044080528 TSecr=1044080528
301423	51.426829	127.0.0.1	127.0.0.1	TCP	68	[TCP Port numbers reused] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044110561 TSecr=0 SACK_PERM=1
301424	51.426872	127.0.0.1	127.0.0.1	TCP	56	[TCP ACKed unseen segment] 9660 → 50995 [ACK] Seq=1 Ack=742040519 Win=407872 Len=0 TSval=1044110561 TSecr=1044080528
301425	51.426875	127.0.0.1	127.0.0.1	TCP	44	50995 → 9660 [RST] Seq=742040519 Win=0 Len=0
301864	51.540336	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044110661 TSecr=0 SACK_PERM=1
301865	51.540345	127.0.0.1	127.0.0.1	TCP	56	[TCP Dup ACK 301424#1] [TCP ACKed unseen segment] 9660 → 50995 [ACK] Seq=1 Ack=742040519 Win=407872 Len=0 TSval=1044110661 TSecr=1044080528
301866	51.540348	127.0.0.1	127.0.0.1	TCP	44	50995 → 9660 [RST] Seq=742040519 Win=0 Len=0
304350	51.653277	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044110761 TSecr=0 SACK_PERM=1
304351	51.653296	127.0.0.1	127.0.0.1	TCP	56	[TCP Dup ACK 301424#2] 9660 → 50995 [ACK] Seq=1 Ack=742040519 Win=407872 Len=0 TSval=1044110761 TSecr=1044080528
304352	51.653299	127.0.0.1	127.0.0.1	TCP	44	50995 → 9660 [RST] Seq=742040519 Win=0 Len=0
305542	51.777827	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044110861 TSecr=0 SACK_PERM=1
305543	51.777833	127.0.0.1	127.0.0.1	TCP	56	[TCP Dup ACK 301424#3] 9660 → 50995 [ACK] Seq=1 Ack=742040519 Win=407872 Len=0 TSval=1044110861 TSecr=1044080528
305544	51.777835	127.0.0.1	127.0.0.1	TCP	44	50995 → 9660 [RST] Seq=742040519 Win=0 Len=0
307635	51.887968	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044110961 TSecr=0 SACK_PERM=1
308396	52.004167	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111061 TSecr=0 SACK_PERM=1
309063	52.228125	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] 50995 → 9660 [SYN] Seq=0 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111261 TSecr=0 SACK_PERM=1
309229	52.245455	127.0.0.1	127.0.0.1	TCP	68	[TCP Previous segment not captured] [TCP Port numbers reused] 9660 → 50995 [SYN, ACK] Seq=894251666 Ack=1 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111276 TSecr=1044110961 SACK_PERM=1
310047	52.253932	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] [TCP Port numbers reused] 9660 → 50995 [SYN, ACK] Seq=4145624494 Ack=1 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111284 TSecr=1044111061 SACK_PERM=1
310879	52.276287	127.0.0.1	127.0.0.1	TCP	68	[TCP Previous segment not captured] [TCP Port numbers reused] 9660 → 50995 [SYN, ACK] Seq=1675207243 Ack=1 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111306 TSecr=1044111261 SACK_PERM=1
311016	52.277571	127.0.0.1	127.0.0.1	TCP	56	[TCP ACKed unseen segment] 50995 → 9660 [ACK] Seq=1 Ack=894251667 Win=408256 Len=0 TSval=1044111307 TSecr=1044111276
311648	52.279550	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [ACK] Seq=1 Ack=894251667 Win=408256 Len=0 TSval=1044111309 TSecr=1044111276
312477	52.281523	127.0.0.1	127.0.0.1	TCP	56	50995 → 9660 [ACK] Seq=1 Ack=894251667 Win=408256 Len=0 TSval=1044111311 TSecr=1044111276
312575	52.281872	127.0.0.1	127.0.0.1	TCP	44	9660 → 50995 [RST] Seq=894251667 Win=0 Len=0
313024	52.283269	127.0.0.1	127.0.0.1	TCP	44	9660 → 50995 [RST] Seq=894251667 Win=0 Len=0
314088	52.383405	127.0.0.1	127.0.0.1	TCP	68	[TCP Retransmission] [TCP Port numbers reused] 9660 → 50995 [SYN, ACK] Seq=1675207243 Ack=1 Win=65535 Len=0 MSS=16344 WS=64 TSval=1044111406 TSecr=1044111261 SACK_PERM=1
314104	52.383434	127.0.0.1	127.0.0.1	TCP	44	50995 → 9660 [RST] Seq=1 Win=0 Len=0
```

A normal HTTP conversation can be seen from operation numbers 58968-60143. The first problem clue can be seen at operation 301423. Note that this operation is some 35 seconds after the HTTP conversation finished. In itself, this is abnormal, as a correct conversation close sequence can be seen between operations 59702 and 60143. So some questions resulted from this analysis:
- Under what conditions are port numbers reported as needing to be reused?
- Why does it appear that ACKs are being given to unseen packets?
- Why are forceable resets being issued (RST)? Normal close sequence is with FIN/ACK.

Ephemeral port reuse is something that can become a resource bottleneck, as there are a limited number of TCP ports, and the number available vary by operating system and network device (in the event of proxies or load balancers). In particular, used ports enter a TIME_WAIT status after use to account for stray packets due to network latency. These ports are unavailable (unless special flags are used) until a timeout expires. This timeout is usually configurable and can span anywhere from 15 seconds to 4 minutes. More information can be found [here][4]. Investigating ephemeral port use seemed logical and fairly easy to perform.

The number of ports in a TIME_WAIT status can be counted using a simplistic status counter, `watch -n 0.5 "netstat -n | grep TIME_WAIT | wc -l"`. This will sum all ports in a TIME_WAIT status every 500 ms. The [number of ports][5] available on the test system (OSX 10.15.7) is 16383. Note that the counter mentioned counts ports without regard to  source IP, source port, destination IP, or destination port. Some port strategies set their limit on this basis. In the OSX case, the port selection strategy is global, and thus the counting strategy mentioned is correct.

Finally, the amount of time to wait is defined in the relevant RFC, however operating systems take liberty with this value. Since the target systems used by Flare are Linux-based, it was decided to adjust the timeout on the test system (OSX with a timeout of 15 seconds) to a timeout typical of Linux systems of 60 seconds. This was done with `sudo sysctl -w net.inet.tcp.msl=60000`.

## Initial Run Results
Four runs were performed, each of 5 minute duration, with a fresh restart of the validator using the startup script mentioned above. The number of simultaneous contracts/addresses was varied for each run.

|Contracts |Max Waits  |Max CPU Util|TPS|Avg ms|Max ms|Min ms|
|--- |--- |--- |--- |--- |--- |--- |
|10|3624|20%|3.15|3088|8017|35|
|40|16340|20%|16.62|2354|8395|43|
|100|16358|40%|22.2|3698|80232|126|
|200|16359|40%|22.41|5911|115089|182|

### Initial Run Observations and Analysis
TPS scaled nearly linearly between runs of 10 vs. 40 contracts, going from 3.15 to 16.62 TPS, while average, max, and minimum transaction times were nearly uniform. CPU utilization was relatively consistent and fairly low at around 20%.

Between runs of 40 to 100 contracts, TPS did not scale linearly and max transaction times increased by an order of magnitude. While CPU utilization nearly doubled, it was observed there were periods of time (2-15 seconds) where there was minimal CPU usage. Note that the maximum number of waits observed is hovering close to the theoretical maximum of 16383.

Between runs of 100 and 200 contracts, TPS clearly reached its maximum under this configuration, with transaction times ever increasing.

## TIME_WAIT Adjustment
In order to test the theory of ephemeral port exhaustion being the source of the resource bottleneck, the timeout was changed from 60 seconds to 2 seconds with `sudo sysctl -w net.inet.tcp.msl=2000`.

### Run Results With 2 Second TIME_WAIT Timeout

|Contracts |Max Waits  |Max CPU Util|TPS|Avg ms|Max ms|Min ms|
|--- |--- |--- |--- |--- |--- |--- |
|40|1053|20%|16.9|2315|4513|50|
|100|2672|40%|45.2|2158|7513|88|
|200|5606|60%|102.7|1884|6908|221|
|500|8439|70%|106.49|4398|8848|2646|
|1000|4000|60%|90.18|9633|16613|5964|

### Observations and Analysis
It is clear that resolving the TCP ephemeral port bottleneck, with a TIME_WAIT = 2 seconds, resulted in TPS throughput to increase linearly to 200 concurrent contracts, with a 4-5x throughput increase over the original test with TIME_WAIT = 60 seconds, while also observing a drastic reduction in transaction times (in particular max time falling by 16x). Beyond 200 concurrent contracts, it is clear that there exists an additional bottleneck that prevents CPU utilization from being saturated.

One theory is that the stress application, while an asynchronous application, may be reaching its limit with its ability to service the number of simultaneous promises resolving (which it does so one at a time as Promise.race returns a resolved promise from the group of contracts awaiting for the increment method to finish). While it may be possible to alter the stress application to more efficiently handle this condition, instead it was decided to run two stress applications simultaneously with a lower load count.

## 2 Stress Applications
Two stress applications were run for a total of five minutes. Each stress application was run with 250 contracts, and an offset of 250 addresses was given to the second stress application such that it would use different source addresses, eliminating nonce conflicts.

### Run Results With 2 Second TIME_WAIT Timeout and 2 Stress Applications

|Contracts |Max Waits  |Max CPU Util|TPS|Avg ms|Max ms|Min ms|
|--- |--- |--- |--- |--- |--- |--- |
|500|4712|75%|123.3|3820|8061|1689|

### Observations and Analysis
TPS increased from 106 to 123 between the the single stress and dual stress application runs of 500 concurrent contracts. While an improvement, which does indicate some deficiency in the single stress application approach at these levels of concurrency, it does not fully explain why throughput did not increase to something over 200 TPS, which would be expected since the CPUs are still not fully saturated. More investigation is required.

## Implications For Other Environments
The test setup is contrived and simple. One should ask as to the applicability of this limitation being relevant or witnessed in other more complex configurations. One substantial difference could be with the use of the websocket protocol vs. HTTP, as TCP port churn with websockets is theoretically lower. There is a [blog][6] that describes TCP ephemeral port exhaustion using the websocket protocol, and so there still exists the possibility.

## Strategies For Remediation
The path followed of reducing the TIME_WAIT value, while a valid approach to testing whether throughput could be increased, is not a recommended approach in a production environment, as it opens up flaws in the TCP protocol of not properly handling late packets due to network latency. Several other alternatives are outlined in the reference materials but those techniques have not yet been tried within the context of the Avalanche application.

## Conclusion and Next Steps
1. Monitor TCP port TIME_WAIT counts within the stress test environment, at the API, bootstrap, load balancing, and core node level. Counts over 10000 should be treated as suspect requiring more analysis and potential attention.
2. If more attention is required, remediation techniques should be analyzed and tried.
3. Perform further investigation of throughput bottlenecks witnessed between the 200-500 simultaneous contract level.

[1]: ./scripts/stress.js
[2]: ./contracts/Counter.sol
[3]: https://www.wireshark.org/
[4]: http://www.serverframework.com/asynchronousevents/2011/01/time-wait-and-its-design-implications-for-protocols-and-scalable-servers.html
[5]: https://dataplane.org/ephemeralports.html
[6]: https://making.pusher.com/ephemeral-port-exhaustion-and-how-to-avoid-it/