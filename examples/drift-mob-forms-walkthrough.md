# Drift Mob Forms Vault Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 184 | ship |
| stress | sync drift | 111 | watch |
| edge | local state | 145 | ship |
| recovery | conflict cost | 192 | ship |
| stale | form pressure | 258 | ship |

Start with `stale` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `form pressure` against `sync drift`, not the raw score alone.
