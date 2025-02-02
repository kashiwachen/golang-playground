# Launch the Project

## 20250202

1. Declare a `main` _package_ to group the group functions under `hello-world` directory.
1. Declare a `main` function, which executes by default when we run the `main` package.

   ```bash
   # /hello-world
   go run .
   ```

1. Call code from external package, `quote`.

   1. Import `rsc.io/quote` in code, and call the following command to:
      1. Add the new module as requirement.
      1. Add into `go.sum` for use in authenticating the module, through computing and verifying the hash.

   ```bash
   # /Anywhere under the repo
   go mod tidy
   ```

1. yes

## 20250127

1. Start to launch project with `go mod init {project_name}`, which would help create a `go,mod` file, which represent the start of a `Go` **module**.

   1. And we will track the dependency of the project in `go.mod`.
   1. P.S. We are using `go-by-test` as `project_name`,
