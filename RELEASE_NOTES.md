# Logger Go Release Notes

## v1.2.6

Lowers the module compatibility floor from Go `1.22` to Go `1.18`, the first
toolchain that provides the `any` alias used by the module's consumer smoke.
The package has no external Go module dependencies. Source, package, and sample
CI cover both Go `1.18.10` and the current stable toolchain.

This patch keeps logger connector/native version `1.2.1`, native generation
`1.2.1+f50756ebff0d`, the public Go API, logger ABI, lifecycle, bounded queue,
pressure behavior, sinks, and all five native payloads unchanged. Go `1.18` is
a compatibility floor, not a recommendation to run an unsupported toolchain
in production; use a currently supported Go release for production builds.
The unchanged Linux payload requires glibc `2.34` or newer and self-reports an
`unknown` git commit. Its generation and source provenance remain pinned by the
native manifest and verified payload digests.
