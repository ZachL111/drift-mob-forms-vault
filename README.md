# drift-mob-forms-vault

`drift-mob-forms-vault` keeps a focused Go implementation around mobile workflows. The project goal is to create a Go reference implementation for forms workflows, centered on constraint solving, bounded scenario files, and conflict explanations.

## Why I Keep It Small

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how form pressure and local state should influence a review result.

## Drift Mob Forms Vault Review Notes

`stale` and `stress` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Included Behavior

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/drift-mob-forms-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `sync drift`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Internal Model

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The Go addition stays small enough to inspect in one sitting.

## Try It Locally

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Validation

The check exercises the source code and the review fixture. `stale` is the high score at 258; `stress` is the low score at 111.

## Scope

The repository is intentionally scoped to local checks. I would expand it by adding adversarial fixtures before adding features.
