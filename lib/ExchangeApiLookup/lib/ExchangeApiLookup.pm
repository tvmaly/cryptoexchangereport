package ExchangeApiLookup;
use Moo;
use HTTP::Request;
use LWP::UserAgent;
use JSON::Parse qw/ json_file_to_perl /;
use strict;
use warnings;
use feature 'say';

has user_agent => ( is => 'lazy' );

sub _build_user_agent {
    return LWP::UserAgent->new();
}



sub commands {
    my $self = shift;
    say "Commands:\n".
        "   - exchanges\n";

}

sub request {
    my $self = shift;
    return HTTP::Request->new( GET => 'https://bittrex.com/api/v1.1/public/getmarkets' );
}

sub send_request {
    my $self = shift;
    my $request = $self->request();
    my $ua = LWP::UserAgent->new();
    return $ua->request($request);
}




1;
