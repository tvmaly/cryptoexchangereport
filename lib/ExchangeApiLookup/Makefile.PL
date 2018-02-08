use ExtUtils::MakeMaker;
use File::ShareDir::Install;

install_share dist => 'share';

WriteMakefile(
    NAME => "ExchangeApiLookup",
    VERSION_FROM => "lib/ExchangeApiLookup.pm",
    EXE_FILES => [
        'scripts/lookup_exchange_api'
    ],
	PREREQ_PM => {
		'Role::Tiny' => 0,
		'Moo' => 0,
		'HTTP::Request' => 0,
		'LWP::UserAgent' => 0,
		'JSON::Parse' => 0,
		'Package::Variant' => 0,
		'Path::Class' => 0,
		'File::ShareDir::Install' => 0,
		'File::Share' => 0,
	}
);

package MY;
use File::ShareDir::Install 'postamble';
