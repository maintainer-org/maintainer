- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-08-03
- RFC PR: [maintainer-org/maintainer#44](https://github.com/maintainer-org/maintainer/pull/44)
- Tracking Issue: [maintainer-org/maintainer#0](https://github.com/maintainer-org/maintainer/issues/0)

# RFC-44: Use git-chglog

## Background

We used to use [github-changelog-generator] for changelog generation.

But [github-changelog-generator] is written in ruby, and users need to install it first:

```shell
$ gem install github_changelog_generator
```

## Proposal

So I propose to use [git-chglog] to replace the dependence on [github-changelog-generator].

## Rationale

N/A

## Compatibility

The template could be incompatible. And we have to set up a migration plan.

It's hard to migrate via coding, maybe we can have the migrate plan in document.

## Implementation

N/A

[github-changelog-generator]: https://github.com/github-changelog-generator/github-changelog-generator
[git-chglog]: https://github.com/git-chglog/git-chglog
