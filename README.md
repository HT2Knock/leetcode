# leetcode

Go solutions to LeetCode problems, kept for practice and for interview prep.
Two older solutions are in TypeScript.

## Layout

- `<number>.<slug>.go`: one solution per problem, named after the problem number
  and the URL slug, for example `1.two-sum.go`. Each file is `package leetcode`,
  and the solution sits between `// @leet start` and `// @leet end`. LeetCode
  uses those markers when you paste the code back into its editor.
- `interview/`: preparation material. `progress.md` tracks topics,
  `results.tsv` records mock scores, and `patterns/`, `mocks/`, and `designs/`
  hold one template each.

## Run a solution

There is no local test harness. A solution file has no `main` and there is no
`go.mod` in the repo, so LeetCode compiles and runs each file in its own
environment. To work on a problem, open the file and edit it, then paste the
marked block into the LeetCode editor for that problem.

## Interview notes

`interview/progress.md` is the status page: which patterns are done, which are
weak. `interview/results.tsv` is the score log from mock interviews. Update both
after each session, so the next session starts from the real state.
