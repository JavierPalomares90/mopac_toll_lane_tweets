package fetcher

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
Fetch the current toll fare from https://mopac-fare.mroms.us/HistoricalFare/
**/

func getRequestData() ([]byte, error) {
	// Add the current time as a parameter to post
	// Must have format MM/DD/YYYY HH:mm
	currTime := time.Now().Format("2006-01-02 15:04")

	params := map[string]string{"starttime": currTime}
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return paramsJSON, nil
}

func FetchCurrentFare() error {
	url := "https://mopac-fare.mroms.us/HistoricalFare/"

	// Get the data to send in the request to get the current fair
	requestData, err := getRequestData()
	if err != nil {
		log.Fatal(err)
		return err
	}
	// send the post request to get the fair
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		log.Fatal(err)
		return err
	}

	// defer closing the response until we're done
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	parseResponseBody(body)
	return nil
}

func parseResponseBody(body []byte) {
    s := string(body)
    log.Print(s)
}

/**
$(function() {
        hisra_current_fetch();
    });

    var url = 'https://mopac-fare.mroms.us/HistoricalFare/';
    var tpn_map = {'cvz':['CVZ to 183', 'CVZ to Parmer'] ,'farwest':['2222 to Parmer'], 'parmer':['Parmer to 2222', 'Parmer to 5th/CVZ'], '2222':['2222 to 5th/CVZ']};
    var direction_map = {'cvz':'nb', 'farwest':'nb', 'parmer':'sb', '2222':'sb'};

    function hisra_current_format_date(dateInfo) {
        function pad(n, width, z) {
            z = z || '0';
            n = n + '';
            return n.length >= width ? n : new Array(width - n.length + 1).join(z) + n;
        }

        if (dateInfo) {
            var fractionalSecond = Math.round(dateInfo.fractionalSecond * 1000) / 1000,
                seconds = pad(dateInfo.second, 2) + '' + fractionalSecond;
            return pad(dateInfo.month, 2)+'/'+pad(dateInfo.day, 2)+'/'+pad(dateInfo.year, 4)+' '+
                pad(dateInfo.hour, 2)+':'+pad(dateInfo.minute, 2)+':'+ pad(dateInfo.second, 2);
        } else return '00/00/0000 00:00:00';

    }

    function hisra_current_format_money(rate) {
        var moola = rate.toFixed(2);
        return '$ '+moola;
    }

    function hisra_current_create_signs(rates) {
        $.each(tpn_map, function(tpn, signs) {
            if ( rates == undefined ) {
                $( '#hisra_sign_' + tpn  ).removeClass();
                $( '#hisra_sign_' + tpn  ).html("<h2>No Sign Information Available for " + tpn + "</h2>");
            }
            else {
                $( '#hisra_sign_' + tpn  ).removeClass();
                $( '#hisra_sign_' + tpn  ).html("");
                $( '#hisra_sign_' + tpn  ).addClass('sign-' + signs.length + '-' + direction_map[tpn]);
                $.each(signs, function(index, sign) {
                    $.each(rates, function(rate) {
                        if ( rates[rate].tpn == sign) {
                            $( '#hisra_sign_' + tpn  ).append('<div style="color: ' + ((rates[rate].rate == 'CLOSED') ? 'red' : 'white') + ';" class="price-box price-box-' + (index + 1) + '">' + rates[rate].rate + '</div>');
                            $( '#hisra_sign_' + tpn  ).append('<div style="color: ' + ((rates[rate].pbm == 'CLOSED') ? 'red' : 'white') + ';" class="price-box-pbm price-box-pbm-' + (index + 1) + '">' + rates[rate].pbm + '</div>');
                        }
                    });
                });
            }
        });
    }

    function hisra_current_fetch() {
        params = {};
        params.starttime = moment(new Date()).format('MM/DD/YYYY HH:mm');

        $.ajax({
            url: 'https://mopac-fare.mroms.us/HistoricalFare/ViewHistoricalFare',
            type: "POST",
            data: params,
            dataType: 'json',
            success: function(data){
                data = data.map(function(datum) {
                    var clean = {};
                    clean.startFormat = hisra_current_format_date(datum.startTimeStamp);
                    clean.tpn = datum.tollingPointName.slice(9, datum.tollingPointName.length);
                    clean.mode = datum.tripMode || 'O';
                    if ( clean.mode == 'I') {
                        clean.rate = 'CLOSED';
                        clean.pbm = 'CLOSED';
                    } else {
                        clean.rate = hisra_current_format_money(datum.tripRate);
                        clean.pbm = hisra_current_format_money(datum.pbmRate || 0.00);
                    }
                    return clean;
                });
                hisra_current_create_signs(data);
            },
            error: function(data) {
                $('#hisra-current-error').css('display', 'block');
            }
        });
    }
**/
