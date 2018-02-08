package Exchanges;

use Package::Variant importing => ['Moo::Role'],
    subs => [ qw( has with ) ];

sub make_variant {
    my ($class, $target_class, $exchanges) = @_;
    use feature 'say';
    use Data::Dumper;

    install 'exchanges' => sub { 
        my $self = shift;
        foreach my $key (keys %$exchanges) {
            say "$key";
        }
        say "\n";
    };

	sub recursively_add_methods {
	    my $name = shift;
	    my $hashmap = shift;
	
	    foreach my $key (keys %$hashmap) {
            my $method_name = defined $name
                ? $name . "_" . $key
                : $key;

	        install $method_name => sub {
	            my $self = shift;
	            return $hashmap->{$key};
	        };

	        recursively_add_methods($method_name, $hashmap->{$key}) if ref $hashmap->{$key} eq 'HASH';
	    }

	}

    recursively_add_methods(undef, $exchanges);

}


1;
