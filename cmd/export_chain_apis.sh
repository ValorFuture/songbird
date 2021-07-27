#!/bin/bash
declare -a SUPPORTED_CHAINS=(LTC XRP)

test_api () {
    if [[ $1 == LTC ]]; then
        if [[ $(curl -sd '{"jsonrpc": "1.0", "id": "curltest", "method": "getnetworkinfo", "params": []}' -u $3:$4 -H 'content-type: text/plain;' $2 | jq .result.subversion) == "" ]]; then
            echo "Invalid LTC API"
            exit;
        fi
    elif [[ $1 == XRP ]]; then
        if [[ $(curl -sH 'Content-Type: application/json' -d '{"method":"server_info","params":[{}]}' -u $3:$4 $2 | jq .result.info.build_version) == "" ]]; then
            echo "Invalid XRP API"
            exit;
        fi
    fi
}

remove_quotes () {
    result="${1%\"}"
    echo "${result#\"}"
}

for CURR_CHAIN in "${SUPPORTED_CHAINS[@]}"
do
    CURR_API_EXPORT=${CURR_CHAIN}_APIs
    CHAIN_APIs=$(cat $1 | jq .$CURR_CHAIN)
    for ((i=0;i<$(echo $CHAIN_APIs | jq '. | length');i++)); 
    do
        CURR_API=$(remove_quotes $(echo $CHAIN_APIs | jq .[$i].api))
        CURR_U=$(remove_quotes $(echo $CHAIN_APIs | jq .[$i].u))
        CURR_P=$(remove_quotes $(echo $CHAIN_APIs | jq .[$i].p))
        echo "Testing API: " $CURR_API
        test_api $CURR_CHAIN $CURR_API $CURR_U $CURR_P
        if [[ $i -eq 0 ]]; then
            declare ${CURR_CHAIN}_APIs=$CURR_API
        else
            declare ${CURR_CHAIN}_APIs=${!CURR_API_EXPORT},$CURR_API
        fi
        CURR_API_HASH=$(printf $CURR_API | sha256sum)
        CURR_API_CHECKSUM=${CURR_API_HASH::8}
        CURR_AUTH_U_EXPORT=${CURR_CHAIN}_U_$CURR_API_CHECKSUM
        CURR_AUTH_P_EXPORT=${CURR_CHAIN}_P_$CURR_API_CHECKSUM
        declare ${CURR_CHAIN}_U_$CURR_API_CHECKSUM=$CURR_U
        declare ${CURR_CHAIN}_P_$CURR_API_CHECKSUM=$CURR_P
        export ${CURR_AUTH_U_EXPORT}=${!CURR_AUTH_U_EXPORT}
        export ${CURR_AUTH_P_EXPORT}=${!CURR_AUTH_P_EXPORT}
    done
    export ${CURR_API_EXPORT}=${!CURR_API_EXPORT}
done