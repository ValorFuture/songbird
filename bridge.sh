PORT=$((8080+$1))
while true; do 
	nohup $(sleep 5; curl -s http://localhost:$PORT/fxrp) >& /dev/null &
	if ! lsof -Pi :$PORT -sTCP:LISTEN -t >/dev/null ; then
	    node fxrp $1
	else
		echo "System already activated."
	fi;
	sleep 60
done