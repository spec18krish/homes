module homes.co.nz

go 1.19

replace homes.co.nz/property/valuations => ./property
replace homes.co.nz/helper => ./helper

require homes.co.nz/property/valuations v0.0.0-00010101000000-000000000000 // indirect
require homes.co.nz/helper v0.0.0-00010101000000-000000000000 // indirect
