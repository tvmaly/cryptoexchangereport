my @exchanges = 
qw(
    bitfinex.com
    bithumb.com
    bitstamp.com
    bittrex.com
    coinone.co.kr
    gdax.com
    gemini.com
    hitbtc.com
    huobi.pro
    kraken.com
    liqui.io
    okex.com
    poloniex.com
);

my @data_types = qw( order_book ticker last_trades );

foreach my $exchange_dir (@exchanges) {

    my $exchange;
    if ($exchange_dir =~ /^(\w+)\./) {
        $exchange = $1;
    } else {
        next;
    }

    foreach my $data_type (@data_types) {
        
        my $file_name = $exchange_dir . "/" . $exchange . "_$data_type";

        print "output $file_name $exchange $data_type\n";
    }
}
