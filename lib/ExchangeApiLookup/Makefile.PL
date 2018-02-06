use ExtUtils::MakeMaker;

WriteMakefile(
    NAME => "ExchangeApiLookup",
    VERSION_FROM => "lib/ExchangeApiLookup.pm",
    EXE_FILES => {
        'scripts/lookup_exchange_api'
    }
);
