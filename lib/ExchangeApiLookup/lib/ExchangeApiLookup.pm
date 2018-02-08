package ExchangeApiLookup;
use strict;
use warnings;
use feature 'say';
use version;
our $VERSION = "1.0.0";

use Moo;
use HTTP::Request;
use LWP::UserAgent;

has user_agent => ( is => 'lazy' );

sub _build_user_agent {
    return LWP::UserAgent->new();
}

sub commands {
    my $self = shift;
    say "Commands:\n".
        "   - exchanges\n".
        "   - request <exchange> <data_type> <symbol>(if symbol is empty, default symbol will be used)\n".
        "   - output <filename> <exchange> <data_type> <symbol>(if symbol is empty, default symbol will be used)\n";
}

sub output {
    my $self = shift;
    my $file = shift;

    my $response = $self->request(@_);

    open(my $fh, '>', $file) or die "Cannot open $file\n";

    print $fh $response;
}
    
sub request {
    my $self = shift;

    my $url = $self->get_url(@_);

    return $self->user_agent->request(HTTP::Request->new( GET => $url ))->content;
}

sub get_url {
    my $self = shift;
    my ($exchange, $data_type, $symbol) = @_;

    $exchange .= "_api_endpoints";

    if (not defined $symbol) {
        my $default_pair = $exchange."_pair_sample";
        $symbol = $self->$default_pair;
    }

    my $url_method = join '_', ($exchange, $data_type);
    my $url = sprintf($self->$url_method, $symbol);

    return $url;
}

1;
