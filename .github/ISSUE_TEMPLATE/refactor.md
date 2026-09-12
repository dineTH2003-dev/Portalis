---
name: Refactor proposal
about: Propose code cleanup or architectural refactoring
title: '[REFACTOR] <Short Description>'
labels: 'type:refactor'
assignees: ''
---

## Current Problem
Describe what in the existing codebase is duplicated, poorly structured, or difficult to maintain. (Adhere to Rule 1: Do not rewrite working code without a demonstrated reason).

## Proposed Change
What files will be changed, moved, or restructured?

## Architectural Impact & Rationale
How does this change align with the architecture specification (`docs/Architecture.md`) and the separation of control plane and data plane?

## Behavior Preservation
Confirm that this refactoring will NOT alter existing public interfaces, protocol schemas, or external behavior.

## Tests & Verification
How will you verify that behavior is strictly preserved?
