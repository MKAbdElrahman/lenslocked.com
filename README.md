# lenslocked.com


A photo sharing website written in Go following the [usegolang.com](https://www.usegolang.com) tutorial.

![CI](https://github.com/MKAbdElrahman/lenslocked.com/actions/workflows/pipeline.yml/badge.svg)


## Getting Started

Developmnet expected to be done on a Linux machine. If you are using Windows, you can use the [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/install-win10) to run the project.

The project uses `Taskfiles` to manage the development tasks. You can find the installation instructions [here](https://taskfile.dev/installation/).

You can run the following tasks to get started:

- `task install-go`: Installs Go
- `task test`: Runs the tests
- `task run`: Runs the application
- `task build`: Builds the application

You can also run `task -a` to see all the available tasks.

All the tasks are defined in `Taskfile.yml`.

The CI pipeline is defined in `.github/workflows/pipeline.yml`. It runs the tests, builds the application and uploads the artifacts to GitHub releases. The pipeline uses `task` to run the tasks.
## Design Choices
- Routing: [gorilla/mux](https://github.com/gorilla/mux)


## Helpful Resources
* [Shipping Go](https://www.manning.com/books/shipping-go)
* [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests)

## Acknowledgements

I would also like to thank [Jon Calhoun](https://www.calhoun.io) for his [Go Web Development](https://www.usegolang.com) course which helped me a lot in understanding the concepts of web development in Go.


