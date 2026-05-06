# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 184, lane `ship`
- `stress`: `sync drift`, score 111, lane `watch`
- `edge`: `local state`, score 145, lane `ship`
- `recovery`: `conflict cost`, score 192, lane `ship`
- `stale`: `form pressure`, score 258, lane `ship`

## Note

The repository should be understandable without pretending it is larger than it is.
